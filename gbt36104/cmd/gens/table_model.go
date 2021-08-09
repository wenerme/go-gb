package gens

type TableMetaModelIndex struct {
	Name string
}

type TableMetaModelColumn struct {
	Name       string
	Type       string
	Comment    string
	PrimaryKey bool
	Unique     bool
	NotNull    bool
}

type TableMetaModel struct {
	Name    string
	Comment string
	Columns []*TableMetaModelColumn
	Indexes []*TableMetaModelIndex
}
