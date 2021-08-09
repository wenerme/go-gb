package gens

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type contextKey string

const (
	GeneratorKey = contextKey("Generator")
	ModelKey     = contextKey("model")
	TemplateKey  = contextKey("template")
)

func (k contextKey) String() string {
	return fmt.Sprintf("gens.contextKey(%s)", string(k))
}

type IsTemplate interface {
	Template() Template
}

type Generator struct {
	Files     []File
	Template  *template.Template
	Templates []IsTemplate
	Debug     bool
	Context   context.Context
	Formatter func(file *File) error
	Loader    func() error
}

type Template struct {
	Name           string // template name.
	Filename       string
	FilenameFormat func(ctx context.Context) string
	Skip           func(ctx context.Context) bool
}

func (t Template) Template() Template {
	return t
}

type WriteConfig struct {
	Target string
	DryRun bool
}

func (ge *Generator) Write(s WriteConfig) error {
	if len(ge.Files) == 0 {
		if err := ge.Generate(); err != nil {
			return err
		}
	}
	for _, v := range ge.Files {
		o := filepath.Join(s.Target, v.Name)
		// log.Println("write ", v.Name, " to ", o)
		if !s.DryRun {
			if err := os.WriteFile(o, v.Content, 0o600); err != nil {
				return err
			}
		} else {
			log.Println(o, "\n", strings.TrimSpace(string(v.Content)))
		}
	}
	return nil
}

func (ge *Generator) Generate(items ...interface{}) error {
	if ge.Debug {
		log.Println("generating")
	}
	if ge.Loader != nil {
		if ge.Debug {
			log.Println("generator loading")
		}
		if err := ge.Loader(); err != nil {
			return err
		}
	}
	files := ge.Files
	templates := ge.Template
	ctx := ge.Context
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = context.WithValue(ctx, GeneratorKey, ge)

	for _, v := range items {
		ctx := context.WithValue(ctx, ModelKey, v)
		for _, ti := range ge.Templates {
			tmpl := ti.Template()
			ctx := context.WithValue(ctx, TemplateKey, v)
			if tmpl.Skip != nil && tmpl.Skip(ctx) {
				continue
			}

			b := bytes.NewBuffer(nil)
			if err := templates.ExecuteTemplate(b, tmpl.Name, v); err != nil {
				return fmt.Errorf("execute template %q: %w", tmpl.Name, err)
			}
			fn := tmpl.Filename
			if tmpl.FilenameFormat != nil {
				fn = tmpl.FilenameFormat(ctx)
			}
			files = append(files, File{
				Name:    fn,
				Content: b.Bytes(),
			})

			if ge.Debug {
				log.Printf("generate %s to %s", tmpl.Name, fn)
			}
		}
	}

	if ge.Formatter != nil {
		for i := range files {
			if err := ge.Formatter(&files[i]); err != nil {
				return err
			}
		}
	}
	ge.Files = files
	return nil
}
