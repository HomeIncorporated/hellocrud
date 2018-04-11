{{- $tableNameSingular := .Table.Name | singular -}}
{{- $modelName := $tableNameSingular | titleCase -}}
{{- $modelNameCamel := $tableNameSingular | camelCase -}}

// {{$modelName}} is an object to back GraphQL type
type {{$modelName}} struct {
  model dbmodel.{{$modelName}}
  db    boil.Executor
}

// GraphQL friendly getters
{{range $column := .Table.Columns }}
{{- if eq $column.Name "id" }}
// Row{{titleCase $column.Name}} is the {{$modelName}} GUID
func (o {{$modelName}}) Row{{titleCase $column.Name}}() string {
  return fmt.Sprintf("{{$modelName}}%d", o.model.{{titleCase $column.Name}}) // {{$column.Type}}
}

// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() graphql.ID {
  return graphql.ID(fmt.Sprintf("%d", o.model.{{titleCase $column.Name}})) // {{$column.Type}}
}
{{- else if eq $column.Type "[]byte" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() []byte {
  return o.model.{{titleCase $column.Name}} // {{$column.Type}}
}
{{- else if eq $column.Type "bool" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() bool {
  return o.model.{{titleCase $column.Name}} // {{$column.Type}}
}
{{- else if eq $column.Type "float32" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() float64 {
  return float64(o.model.{{titleCase $column.Name}}) // {{$column.Type}}
}
{{- else if eq $column.Type "float64" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() float64 {
  return o.model.{{titleCase $column.Name}} // {{$column.Type}}
}
{{- else if eq $column.Type "int" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() int32 {
  return int32(o.model.{{titleCase $column.Name}}) // {{$column.Type}}
}
{{- else if eq $column.Type "int16" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() int32 {
  return int32(o.model.{{titleCase $column.Name}}) // {{$column.Type}}
}
{{- else if eq $column.Type "int64" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() int32 {
  return int32(o.model.{{titleCase $column.Name}}) // {{$column.Type}}
}
{{- else if eq $column.Type "null.Bool" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() *bool {
  return o.model.{{titleCase $column.Name}}.Ptr() // {{$column.Type}}
}
{{- else if eq $column.Type "null.Byte" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() *byte {
  return o.model.{{titleCase $column.Name}}.Ptr() // {{$column.Type}}
}
{{- else if eq $column.Type "null.Bytes" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() *[]byte {
  return o.model.{{titleCase $column.Name}}.Ptr() // {{$column.Type}}
}
{{- else if eq $column.Type "null.Float64" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() *float64 {
  return o.model.{{titleCase $column.Name}}.Ptr() // {{$column.Type}}
}
{{- else if eq $column.Type "null.Int" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() *int32 {
  if !o.model.{{titleCase $column.Name}}.Valid {
    return nil
  }
  x := int32(o.model.{{titleCase $column.Name}}.Int) // {{$column.Type}}
  return &x
}
{{- else if eq $column.Type "null.Int16" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() *int32 {
  if !o.model.{{titleCase $column.Name}}.Valid {
    return nil
  }
  x := int32(o.model.{{titleCase $column.Name}}.Int16) // {{$column.Type}}
  return &x
}
{{- else if eq $column.Type "null.Int64" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() *int32 {
  if !o.model.{{titleCase $column.Name}}.Valid {
    return nil
  }
  x := int32(o.model.{{titleCase $column.Name}}.Int64) // {{$column.Type}}
  return &x
}
{{- else if eq $column.Type "null.JSON" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() *[]byte {
  return o.model.{{titleCase $column.Name}}.Ptr() // {{$column.Type}}
}
{{- else if eq $column.Type "null.String" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() *string {
  return o.model.{{titleCase $column.Name}}.Ptr() // {{$column.Type}}
}
{{- else if eq $column.Type "null.Time" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() *graphql.Time {
  if o.model.{{titleCase $column.Name}}.Valid {
		return &graphql.Time{Time: o.model.{{titleCase $column.Name}}.Time}
	}
	return nil // {{$column.Type}}
}
{{- else if eq $column.Type "string" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() string {
  return o.model.{{titleCase $column.Name}} // {{$column.Type}}
}
{{- else if eq $column.Type "time.Time" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() graphql.Time {
  return graphql.Time{Time: o.model.{{titleCase $column.Name}}} // {{$column.Type}}
}
{{- else if eq $column.Type "types.Byte" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() string {
  return o.model.{{titleCase $column.Name}}.String() // {{$column.Type}}
}
{{- else if eq $column.Type "types.JSON" }}
// {{titleCase $column.Name}} is the {{$modelName}} {{$column.Name}}
func (o {{$modelName}}) {{titleCase $column.Name}}() string {
  return o.model.{{titleCase $column.Name}}.String() // {{$column.Type}}
}
{{- end -}}
{{- end }}
