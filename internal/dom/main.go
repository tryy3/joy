//go:generate go run main.go

package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/joy/internal/dom/curl"
	"github.com/matthewmueller/joy/internal/dom/def"
	"github.com/matthewmueller/joy/internal/dom/generatestructure"
	"github.com/matthewmueller/joy/internal/dom/graph"
	"github.com/matthewmueller/joy/internal/dom/index"
	"github.com/matthewmueller/joy/internal/dom/parser"
	"github.com/matthewmueller/joy/internal/gen"
	"github.com/pkg/errors"
)

var browserAPIURL = "https://rawgit.com/Microsoft/TSJS-lib-generator/master/inputfiles/browser.webidl.xml"
var browserDocsURL = "https://rawgit.com/Microsoft/TSJS-lib-generator/master/inputfiles/comments.json"
var structurePATH = path.Join("github.com", "matthewmueller", "joy", "internal", "dom", "generatestructure", "structure.json")

// manually selected package names for the cliques
var cliqueNames = map[string]string{
	"IntersectionObserver": "intersectionobserver",
	"MutationObserver":     "mutationobserver",
	"MimeType":             "mimetype",
	"Window":               "window",
	"MediaQueryList":       "mediaquery",
	"IDBDatabase":          "idb",
	"VideoTrack":           "avtrack",
	"TextTrack":            "texttrack",
	"MSHTMLWebViewElement": "mswebviewelement",
	"SVGUseElement":        "svguseelement",
	"WebKitFileSystem":     "webkitfilesytem",
	"AudioNode":            "audionode",
	"Element":              "element",
}

func generate(dir string) error {
	structure, err := generatestructure.LoadStructure(path.Join(build.Default.GOPATH, "src", structurePATH))
	if err != nil {
		return errors.Wrapf(err, "error loading structure.json")
	}

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("couldn't get file from runtime")
	}
	dirname := path.Dir(file)

	xml, err := ioutil.ReadFile(path.Join(dirname, "inputs", "browser.webidl.xml"))
	if err != nil {
		return err
	}

	comments, err := curl.JSON(browserDocsURL)
	if err != nil {
		return err
	}

	index, err := parser.Parse(string(xml), comments)
	if err != nil {
		return err
	}

	// idx, err := override.Override(index)
	// if err != nil {
	// 	return errors.Wrapf(err, "error overriding")
	// }
	// index = idx

	g := graph.New()
	for _, parent := range index {
		g.Node(parent)

		deps, err := parent.Dependencies()
		if err != nil {
			return err
		}
		for _, child := range deps {
			g.Edge(parent, child)
		}
	}

	// generate the packages
	cliques := g.Cliques()
	for _, clique := range cliques {
		if len(clique) == 1 {
			name := clique[0].ID()
			pkgname := gen.Lowercase(name)
			index[name].SetPackage(pkgname)
			index[name].SetFile(pkgname)
			continue
		}

		pkgname := ""
		for _, def := range clique {
			if cliqueNames[def.ID()] != "" {
				pkgname = cliqueNames[def.ID()]
				break
			}
		}
		if pkgname == "" {
			var ids []string
			for _, def := range clique {
				ids = append(ids, def.ID())
			}
			return fmt.Errorf("group name not defined for this clique: %s", strings.Join(ids, ", "))
		}

		log.Infof("clique %s", pkgname)
		for _, def := range clique {
			log.Infof("- %s", def.ID())
			name := def.ID()
			filename := gen.Lowercase(name)
			def := index[name]
			def.SetPackage(pkgname)
			def.SetFile(filename)
		}
	}

	sortIndex(structure, index)

	if err = structure.SaveStructure(); err != nil {
		return errors.Wrapf(err, "error saving structure")
	}

	// accurate length
	var defs []def.Definition
	for _, def := range index {
		// only use these as these are our top-level files
		switch def.Kind() {
		case "ENUM", "DICTIONARY", "INTERFACE":
			defs = append(defs, def)
		}
	}

	// write first so we have all the files present
	if err = generateFile(defs, dir); err != nil {
		return err
	}

	// outpath := path.Join(dir, "window")
	// if err := os.MkdirAll(outpath, 0755); err != nil {
	// 	return errors.Wrapf(err, "error mkdir")
	// }

	// output := "package dom\n\n" + strings.Join(codes, "\n\n")
	// formatted, err := gen.Format(string(output))
	// if err != nil {
	// 	if e := ioutil.WriteFile(path.Join(dir, "dom.go"), []byte(output), 0644); e != nil {
	// 		return errors.Wrapf(e, "error writing dom.go")
	// 	}
	// 	return errors.Wrapf(err, "error formatting")
	// }

	// if e := ioutil.WriteFile(path.Join(dir, "dom.go"), []byte(formatted), 0644); e != nil {
	// 	return errors.Wrapf(e, "error writing dom.go")
	// }
	// fmt.Println(formatted)

	// format and link all the packages up
	l := len(defs)
	for i, def := range defs {
		filepath := path.Join(dir, def.GetPackage(), def.GetFile()+".go")

		buf, err := ioutil.ReadFile(filepath)
		if err != nil {
			return errors.Wrapf(err, "error reading file")
		}

		code, err := gen.Format(string(buf))
		if err != nil {
			return errors.Wrapf(err, "error formatting %s", def.ID())
		}

		if err := ioutil.WriteFile(filepath, []byte(code), 0644); err != nil {
			return errors.Wrapf(err, "error writing formatted file")
		}

		log.Infof("formatted %s (%d/%d)", def.ID(), i, l)
	}

	return nil
}

func sortIndex(expected *generatestructure.Structure, index index.Index) {
	// handle all already confirmed moved packages
	for _, i := range index {
		for _, pkg := range expected.Packages {
			for _, file := range pkg.Files {
				if file.Name == i.GetFile()+".go" && pkg.Name != i.GetPackage() && file.Confirmed {
					oldpkg := i.GetPackage()
					i.SetPackage(pkg.Name)
					log.Infof("moved %s from package %s to %s", i.GetFile(), oldpkg, pkg.Name)
				}
			}
		}
	}

	// go through all indexes again and try to move more files without creating circular problems
	for _, i := range index {
		for _, pkg := range expected.Packages {
			for _, file := range pkg.Files {
				if file.Name == i.GetFile()+".go" && pkg.Name != i.GetPackage() {
					if !file.Confirmed {
						oldpkg := i.GetPackage()
						changePackage(index, i, pkg.Name)

						if oldpkg != i.GetPackage() {
							log.Infof("moved and confirmed %s from package %s to %s", i.GetFile(), oldpkg, pkg.Name)
							file.Confirmed = true
						}

						if file.Name == "readablestream.go" {
							log.Infof("%#v - %#v - %#v\n", oldpkg, i.GetPackage(), file)
						}
					}
				}
			}
		}
	}
}

func generateFile(defs []def.Definition, dir string) error {
	e := make(chan error)
	fileCreated := make(chan string)
	l := len(defs)

	// write first so we have all the files present
	for _, d := range defs {
		go func(def def.Definition) {
			code, err := def.Generate()
			if err != nil {
				e <- errors.Wrapf(err, "error generating %s", def.ID())
				return
			}

			// add the package name at the top
			code = "package " + def.GetPackage() + "\n\n" + code
			pkgpath := path.Join(dir, def.GetPackage())
			if err := os.MkdirAll(pkgpath, 0755); err != nil {
				e <- errors.Wrapf(err, "error mkdir")
				return
			}

			if err := ioutil.WriteFile(path.Join(pkgpath, def.GetFile()+".go"), []byte(code), 0644); err != nil {
				e <- errors.Wrapf(err, "error writefile")
				return
			}
			fileCreated <- def.GetFile()
		}(d)
	}

	count := 0
	for {
		count++
		select {
		case err := <-e:
			return err
			//case <-fileCreated:
		case name := <-fileCreated:
			log.Infof("generated %s (%d/%d)", name, count, l)
			if count >= l {
				return nil
			}
		}
	}
}

func contains(list []string, find string) bool {
	for _, s := range list {
		if s == find {
			return true
		}
	}
	return false
}

func changePackage(index index.Index, def def.Definition, pkg string) {
	visited := &[]string{}

	dependencies, err := def.Dependencies()
	if err != nil {
		return
	}

	oldpkg := def.GetPackage()
	def.SetPackage(pkg)

	// Start looping through all dependencies
	for _, dependency := range dependencies {
		// get the dependencies and start checking for circular problems
		deps, err := dependency.Dependencies()
		if err != nil {
			continue
		}
		for _, dep := range deps {
			if followDep(index, dep.GetPackage(), dep, visited) {
				def.SetPackage(oldpkg)
				return
			}
		}
	}

	// Check for all definitions that depend on the changed definition
	/*for _, i := range index {
		// we have already updated this definition
		if i.GetFile() == def.GetFile() {
			continue
		}

		deps, err := i.Dependencies()
		if err != nil {
			return
		}

		// go through all the dependencies
		for _, dep := range deps {
			// check if this definition depends on the definition we updated earlier

			if def.GetFile() == "positionerror" && dep.GetFile() == def.GetFile() {
				fmt.Printf("%s: %#v\n", i.GetFile(), dep.ID())
			}
			if dep.GetFile() == def.GetFile() {
				if i.GetPackage() == pkg {
					i.SetPackage("")
					break
				}
				i.SetPackage(pkg)
				break
			}
		}
	}*/
}

// followDep will try to follow all dependencies and look for any circular dependency issues
func followDep(index index.Index, findpkg string, def def.Definition, visited *[]string) bool {
	// we have already visited this file
	if contains(*visited, def.GetFile()) {
		return false
	}
	*visited = append(*visited, def.GetFile())

	// TODO: Is this really needed since we are looking at all the deps?
	if def.GetPackage() == findpkg {
		return true
	}

	// get dependencies and check if there is any dependencies to check
	deps, err := def.Dependencies()
	if err != nil {
		return false
	}

	for _, dep := range deps {
		// check for circular import issues
		if dep.GetPackage() == findpkg {
			return true
		}

		// follow the dependency
		if followDep(index, findpkg, dep, visited) {
			return true
		}

		// follow all the files in the package
		for _, i := range index {
			if dep.GetPackage() == i.GetPackage() {
				if followDep(index, findpkg, dep, visited) {
					return true
				}
			}
		}
	}

	return false
}

func unique(defs []def.Definition) []def.Definition {
	u := make([]def.Definition, 0, len(defs))
	m := make(map[string]bool)

	for _, def := range defs {
		if _, ok := m[def.ID()]; !ok {
			m[def.ID()] = true
			u = append(u, def)
		}
	}

	return u
}

func main() {
	log.SetHandler(text.New(os.Stderr))

	pwd, err := os.Getwd()
	if err != nil {
		log.WithError(err).Fatalf("error getting cwd")
	}

	// if err := os.RemoveAll(path.Join(pwd, "dom2")); err != nil {
	// 	log.WithError(err).Fatalf("removing dom")
	// }

	if e := generate(path.Join(pwd, "dom")); e != nil {
		log.WithError(e).Fatalf("error generating")
	}
}
