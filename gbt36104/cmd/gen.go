package main

import (
	"embed"
	"flag"
	"log"
	"os"
	"reflect"
	"text/template"

	"github.com/wenerme/go-gb/gbt36104/cmd/gens"

	"github.com/Masterminds/sprig"

	"gopkg.in/yaml.v3"
)

//go:embed tpl/*.tmpl
var templateFS embed.FS

func MustParseTemplates() *template.Template {
	return template.Must(
		template.New("tpl").
			Funcs(sprig.TxtFuncMap()).
			Funcs(template.FuncMap{
				"last": func(x int, a interface{}) bool {
					return x == reflect.ValueOf(a).Len()-1
				},
			}).
			ParseFS(templateFS, "tpl/*.tmpl"),
	)
}

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
	mm := &gens.EntityMetaModel{}
	if err := yaml.Unmarshal(file, mm); err != nil {
		log.Fatalln(err)
	}
	NoError(gens.Normalize(mm))

	g := &gens.Generator{
		Template: MustParseTemplates(),
		Templates: []gens.IsTemplate{
			gens.MetaModelTemplate{
				Name:     "go/model",
				Filename: "model.go",
			},
			gens.MetaModelTemplate{
				Name:     "sql/pg",
				Filename: "model.pg.sql",
			},
		},
		Formatter: gens.GoFormatter,
	}
	NoError(g.Generate(mm))
	NoError(g.Write(gens.WriteConfig{}))
}

func NoError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
