package {{.Model}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
	{{range .Imports}}"{{.}}"{{end}}
)
{{end}}

{{range .Tables}}
type {{Mapper .Name}} struct {
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
{{end}}
}

func ({{.Model}} a*{{.Name}}) SetUpdatedBy(trace string) {
	{{.Model}}.UpdatedBy = trace
}

func ({{.Model}} *{{.Name}}) BeforeInsert() {
	{{.Model}}.CreatedAt = time.Now()
	{{.Model}}.UpdatedAt = time.Now()
	{{.Model}}.CreatedBy = db.UpdatedBy
}

func ({{.Model}} *{{.Name}}) BeforeUpdate() {
	{{.Model}}.UpdatedAt = time.Now()
}

{{end}}
