package gens

import (
	"mvdan.cc/gofumpt/format"
	"strings"

	"golang.org/x/tools/imports"
)

type File struct {
	Name    string
	Content []byte
}

func gofmt(f []byte) (i []byte, err error) {
	return format.Source(f, format.Options{
		ExtraRules: true,
	})
}

func GoFormatter(f *File) (err error) {
	if !strings.HasSuffix(f.Name, ".go") {
		return
	}
	f.Content, err = gofmt(f.Content)
	if err == nil {
		f.Content, err = imports.Process(f.Name, f.Content, nil)
	}
	return
}
