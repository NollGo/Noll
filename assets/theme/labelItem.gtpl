{{define "LabelTemplate"}}
<span {{ template "ColorStyleTemplate" .Color }}>#{{ .Name }}</span>
{{end}}