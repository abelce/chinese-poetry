{{define "operand"}}
switch cond.Operand {
{{if eq .Type "bool"}} 
    {{template "bool" .}}
{{end}}
{{if eq .Type "string"}}
    {{template "string" .}}
{{end}}
{{if isNumber .Type}}
    {{template "number" .}}
{{end}}
}
{{end}}