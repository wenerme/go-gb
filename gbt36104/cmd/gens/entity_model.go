package gens

import (
	"context"
	"fmt"

	"github.com/huandu/xstrings"
	"github.com/jinzhu/inflection"
)

type EntityMetaModelField struct {
	Name    string
	NameZh  string `yaml:"nameZh"`
	Type    string
	SQLType string `yaml:"sqlType"`
	GoType  string `yaml:"goType"`
}

type EntityMetaModel struct {
	Name      string
	TableName string `yaml:"tableName"`
	Fields    []*EntityMetaModelField
}

type MetaModelTemplate struct {
	Name           string // template name.
	Filename       string
	FilenameFormat func(ctx context.Context, mm *EntityMetaModel) string
	Skip           func(ctx context.Context, mm *EntityMetaModel) bool
}

func (mt MetaModelTemplate) Template() Template {
	t := Template{
		Name:           mt.Name,
		Filename:       mt.Filename,
		FilenameFormat: nil,
		Skip: func(ctx context.Context) bool {
			mm, ok := ctx.Value(ModelKey).(*EntityMetaModel)
			if !ok {
				return true
			}
			if mt.Skip != nil {
				return mt.Skip(ctx, mm)
			}
			return false
		},
	}
	if mt.FilenameFormat != nil {
		t.FilenameFormat = func(ctx context.Context) string {
			mm := ctx.Value(ModelKey).(*EntityMetaModel)
			return mt.FilenameFormat(ctx, mm)
		}
	}
	return t
}

func Normalize(mm *EntityMetaModel) error {
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
		if f.GoType == "" {
			return fmt.Errorf("no go type %q", f.Name)
		}
	}
	return nil
}
