{{define "CategoryItemTemplate"}}
<a href="{{ url . }}">
  <span>{{ .EmojiHTML }} {{ .Name }}</span>
  {{ if .Discussions }}
  ({{ .Discussions.TotalCount }})
  {{ end }}
</a>
{{end}}