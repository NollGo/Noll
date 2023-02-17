{{define "LabelItemTemplate"}}
<a href="{{ url . }}" {{ template "ColorStyleTemplate" .Color }}>
  <span>#{{ .Name }}</span>
  {{ if .Discussions }}
  ({{ .Discussions.TotalCount }})
  {{ end }}
</a>
{{end}}