package cli

import (
	"bufio"
	"context"
	"fmt"
	"github.com/andrew-d/go-termutil"
	"github.com/matthewmueller/joy/internal/paths"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/prompt"
	"github.com/matthewmueller/joy/internal/stats"
	"github.com/matthewmueller/store"

	"github.com/pkg/errors"

	apibuild "github.com/matthewmueller/joy/api/build"
	"github.com/matthewmueller/joy/internal/cli/build"
	"github.com/matthewmueller/joy/internal/cli/run"
	"github.com/matthewmueller/joy/internal/cli/serve"
	"github.com/matthewmueller/joy/internal/cli/test"
	"github.com/matthewmueller/joy/internal/cli/upgrade"
	"github.com/matthewmueller/joy/internal/cli/version"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// Run the CLI
func Run(ctx context.Context, ver string) (err error) {
	// setup stats
	defer func() {
		if err := stats.Client.MaybeFlush(100, 1*time.Minute); err != nil {
			log.WithError(err).Errorf("error flushing")
		}
	}()

	stats.Client.Set(map[string]interface{}{
		"os":      runtime.GOOS,
		"arch":    runtime.GOARCH,
		"version": ver,
	})

	cmd := kingpin.New("joy", "Joy – A Delightful Go to Javascript Compiler")
	cmd.Version(ver)

	// setup our local db
	db, err := store.New("joy")
	if err != nil {
		return errors.Wrapf(err, "unable to setup the storage")
	}

	runs, err := stats.Increment(db)
	if err != nil {
		return errors.Wrapf(err, "error incrementing stats")
	}

	if ver != "master" {
		if done, err := prompt.Prompt(db, runs); err != nil || done {
			return errors.Wrapf(err, "prompt error")
		}
	}

	// special case: joy ./main & ./main.go or when someone pipes data to
	stat, _ := os.Stdin.Stat()
	if len(os.Args[1:]) == 1 || (stat.Mode()&os.ModeCharDevice == 0) {
		file := ""
		if len(os.Args[1:]) == 1 {
			file = os.Args[1]
		} else {
			defer removeTempFile()
			p, err := paths.Joy()
			if err != nil {
				return errors.Wrapf(err, "error getting joy's root path")
			}

			goSrc, err := readStreamInput()
			if err != nil {
				return errors.Wrapf(err, "error reading stream input")
			}

			tmpDir := filepath.Join(p, "tmp")
			if err := os.Mkdir(tmpDir, os.ModePerm); err != nil {
				return errors.Wrapf(err, "error creating temp folder")
			}

			file = filepath.Join(tmpDir, "main.go")
			if err := ioutil.WriteFile(file, []byte(goSrc), os.ModePerm); err != nil {
				return errors.Wrapf(err, "error creating temp file")
			}
		}

		if _, err := os.Stat(file); !os.IsNotExist(err) {
			files, err := apibuild.Build(&apibuild.Config{
				Context:  ctx,
				Packages: []string{file},
			})
			if err != nil {
				return errors.Wrapf(err, "error building code")
			} else if len(files) != 1 {
				return fmt.Errorf("joy run expects only 1 main file, but received %d files", len(files))
			}

			fmt.Println(files[0].Source())
			return nil
		}
	}

	// commands
	run.New(ctx, cmd)
	build.New(ctx, cmd)
	serve.New(ctx, cmd)
	test.New(ctx, cmd)
	upgrade.New(ctx, cmd, ver)
	version.New(ctx, cmd, ver)

	// run the command
	_, err = cmd.Parse(os.Args[1:])
	return err
}

func readStreamInput() (string, error) {
	r := bufio.NewReader(os.Stdin)
	str := ""
	buf := make([]byte, 0, 4*1024)
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			return "", err
		}
		// process buf
		str += string(buf)
		if err != nil && err != io.EOF {
			return "", err
		}
	}
	return str, nil
}

func removeTempFile() {
	p, err := paths.Joy()
	if err != nil {
		return
	}

	tmpDir := filepath.Join(p, "tmp")
	err = os.RemoveAll(tmpDir)
}
