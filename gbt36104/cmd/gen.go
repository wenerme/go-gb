package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"text/template"

	"github.com/wenerme/go-gens/gengo"
	"github.com/wenerme/go-gens/gengqls"

	"github.com/huandu/xstrings"
	"github.com/jinzhu/inflection"

	"github.com/wenerme/go-gens/gen"
	"github.com/wenerme/go-gens/models/entm"

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
	mm := &entm.EntityMetaModel{}
	if err := yaml.Unmarshal(file, mm); err != nil {
		log.Fatalln(err)
	}
	NoError(entm.Normalize(mm))

	g := &gen.Generator{
		Template: MustParseTemplates(),
		Templates: []gen.IsTemplate{
			entm.MetaModelTemplate{
				Name:     "go/model",
				Filename: "model.go",
			},
			entm.MetaModelTemplate{
				Name:     "sql/pg",
				Filename: "model.pg.sql",
			},
			entm.MetaModelTemplate{
				Name:     "gqls/type",
				Filename: "model.graphqls",
			},
		},
		Formatter: gen.Formatters(gengo.Format, gengqls.Format),
	}
	NoError(g.Generate(mm))
	NoError(g.Write(gen.WriteConfig{}))
}

func NoError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Normalize(mm *entm.EntityMetaModel) error {
	if mm.TableName == "" {
		mm.TableName = inflection.Plural(xstrings.ToSnakeCase(mm.Name))
	}
	for _, f := range mm.Fields {
		if f.Type == "" {
			f.Type = "string"
		}
		if f.SQLType == "" {
			switch f.Type {
			case "string":
				f.SQLType = "text"
			default:
				f.SQLType = f.Type
			}
		}
		if f.GoType == "" {
			f.GoType = f.Type
		}
		if f.GoType == "float" {
			f.GoType = "float64"
		}
		if f.GoType == "" {
			return fmt.Errorf("no go type %q", f.Name)
		}
	}
	return nil
}
