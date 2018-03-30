package core

import (
	"io/ioutil"
	"log"
	"strings"
)

type Helper interface {
	LoadDepFile(filepath string)
	LatestDepVersion(dep string) string
	FetchDepInfo(dep Dep) []byte
}

type Dep struct {
	Name    string
	Version string
}

type PythonHelper struct {
	BaseUrl string
	Deps    []Dep
}

func (ph *PythonHelper) LoadDepFile(filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Failed to load dep file: %v", err)
	}
	ph.Deps = ParseDepFile(raw)
}

func ParseDepFile(raw []byte) []Dep {
	deps := []Dep{}
	lines := splitLines(string(raw[:len(raw)]))

	for _, line := range lines {
		split := splitDep(line)
		deps = append(deps, Dep{split[0], split[1]})
	}

	return deps
}

func splitLines(s string) []string {
	return strings.Split(s, "\n")
}

func splitDep(s string) []string {
	return strings.Split(s, "==")
}

type StrFn func(s string) string

func compose(fns ...StrFn) StrFn {
	if len(fns) == 2 {
		return func(s string) string {
			return fns[0](fns[1](s))
		}
	}

	newFns := fns[0 : len(fns)-2]
	newFns = append(newFns,
		func(s string) string {
			return fns[len(fns)-2](fns[len(fns)-1](s))
		})

	return compose(newFns...)
}
