package dwarf

import (
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strings"
)

const path = "./data/names.txt"

// Contains all utility functions
// to be used with dwarves at runtime.
type DwarfService struct {
	Names []string
}

func NewService() DwarfService {
	return DwarfService{Names: loadNames()}
}

func (d *DwarfService) RandomName() string {
	return d.Names[rand.Intn(len(d.Names)-1)]
}

// Loads names files, sorts names
// and saves it. Useful for keeping
// order when adding/removing names.
func (d *DwarfService) CleanNames() {
	names := loadNames()
	saveNames(names)
}

func loadNames() []string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var names []string
	for _, name := range strings.Split(string(f), "\r") {
		n := strings.ReplaceAll(name, "\n", "")
		if n == "" {
			continue
		}
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

func saveNames(s []string) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, line := range s {
		_, err = f.WriteString(line + "\r")
		if err != nil {
			panic(err)
		}
	}
}
