{{define "LabelItemTemplate"}}
<a href="/label/{{ .Name }}.html" {{ template "ColorStyleTemplate" .Color }}>
  <span>#{{ .Name }}</span>
  {{ if .Discussions }}
  ({{ .Discussions.TotalCount }})
  {{ end }}
</a>
{{end}}