package generatestructure

import (
	"encoding/json"
	"io/ioutil"
)

type Package struct {
	Name  string  // Package names
	Files []*File // All files in the package
}

type Structure struct {
	path     string
	Packages []*Package
}

func NewStructure(path string) *Structure {
	return &Structure{path: path}
}

func LoadStructure(path string) (*Structure, error) {
	structure := &Structure{path: path}

	packages := []*Package{}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &packages)
	if err != nil {
		return nil, err
	}

	structure.Packages = packages
	return structure, nil
}

func (s *Structure) SaveStructure() error {
	b, err := json.MarshalIndent(s.Packages, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.path, b, 0644)
}

func (s *Structure) GetPackage(pkg string) *Package {
	for _, p := range s.Packages {
		if p.Name == pkg {
			return p
		}
	}
	return nil
}

func (s *Structure) GetFile(file string) *File {
	for _, p := range s.Packages {
		for _, f := range p.Files {
			if f.Name == file {
				return f
			}
		}
	}
	return nil
}
