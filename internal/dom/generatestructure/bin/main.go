package main

import (
	"path"
	"go/build"
	"flag"
	"io/ioutil"
	"log"
	"github.com/matthewmueller/joy/internal/dom/generatestructure"
	"path/filepath"
)

 func getFiles(path string) (map[string][]string, error) {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	outputPackages := map[string][]string{}

	for _, dir := range dirs {
		outputFiles := []string{}

		if dir.IsDir() {
			files, err := ioutil.ReadDir(path + string(filepath.Separator) + dir.Name())
			if err != nil {
				return nil, err
			}
			for _, file := range files {
				if !file.IsDir() {
					outputFiles = append(outputFiles, file.Name())
				}
			}
		}

		outputPackages[dir.Name()] = outputFiles
	}
	return outputPackages, nil
}

var p = flag.String("path", path.Join(build.Default.GOPATH, "src", "github.com", "matthewmueller", "joy"), "Path to joy project")

func main () {
	flag.Parse()

	structureJson := path.Join(*p, "internal", "dom", "generatestructure", "structure.json")

	structure, err := generatestructure.LoadStructure(structureJson)
	if err != nil {
		log.Panic(err)
	}

	packages, err := getFiles(path.Join(*p, "dom"))
	if err != nil {
		log.Panic(err)
	}

	for pkg, files := range packages {
		p := structure.GetPackage(pkg)
		if p == nil {
			continue
		}

		for _, file1 := range files {
			for _, file2 := range p.Files {
				if file1 == file2.Name {
					file2.Confirmed = true
					break
				}
			}
		}
	}

	if err = structure.SaveStructure(); err != nil {
		log.Panic(err)
	}
}