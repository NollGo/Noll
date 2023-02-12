{{define "CategoryItemTemplate"}}
<a href="/category/{{ .Name }}.html">
  <span>{{ .EmojiHTML }} {{ .Name }}</span>
  {{ if .Discussions }}
  ({{ .Discussions.TotalCount }})
  {{ end }}
</a>
{{end}}