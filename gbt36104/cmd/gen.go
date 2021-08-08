package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"

	"github.com/Masterminds/sprig"

	"gopkg.in/yaml.v3"
)

func main() {
	metaFile := ""
	flag.StringVar(&metaFile, "i", "model.yaml", "metafile")
	flag.Parsed()
	if metaFile == "" {
		log.Fatalln("no model file")
	}
	file, err := os.ReadFile(metaFile)
	if err != nil {
		log.Fatalln(err)
	}
	mm := &MetaModel{}
	if err := yaml.Unmarshal(file, mm); err != nil {
		log.Fatalln(err)
	}
	NoError(Normalize(mm))

	tpl := template.New("go")
	tpl.Funcs(sprig.TxtFuncMap())
	_, err = tpl.Parse(`
package gbt36104

type {{.Name}} struct {
  {{- range $k, $v := .Fields}}
  {{$v.Name | title}} {{$v.GoType}} // {{$v.NameZh}} 
  {{- end}}
}
`)
	NoError(err)
	out := &bytes.Buffer{}
	NoError(tpl.Execute(out, mm))
	b := out.Bytes()
	b, err = gofmt(b)
	NoError(err)
	NoError(os.WriteFile("model.go", b, 0o600))
}

func NoError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type MetaModelField struct {
	Name   string
	NameZh string `yaml:"nameZh"`
	Type   string
	GoType string `yaml:"goType"`
}

type MetaModel struct {
	Name   string
	Fields []*MetaModelField
}

func Normalize(mm *MetaModel) error {
	for _, f := range mm.Fields {
		if f.Type == "" {
			f.Type = "string"
		}
		if f.GoType == "" {
			f.GoType = f.Type
		}
		if f.GoType == "" {
			return fmt.Errorf("no go type %q", f.Name)
		}
	}
	return nil
}

func gofmt(f []byte) (i []byte, err error) {
	return format.Source(f)
}
