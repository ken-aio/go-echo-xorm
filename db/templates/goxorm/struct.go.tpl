{{$packName := .Model}}
package {{$packName}}

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

func ({{$packName}} *{{Mapper .Name}}) SetUpdatedBy(trace string) {
	{{$packName}}.UpdatedBy = trace
}

func ({{$packName}} *{{Mapper .Name}}) BeforeInsert() {
	{{$packName}}.CreatedAt = time.Now()
	{{$packName}}.UpdatedAt = time.Now()
	{{$packName}}.CreatedBy = db.UpdatedBy
}

func ({{$packName}} *{{Mapper .Name}}) BeforeUpdate() {
	{{$packName}}.UpdatedAt = time.Now()
}

{{end}}
