{{define "LabelGroupTemplate"}}
<h1>标签({{ .Labels.TotalCount }})</h1>
{{ if .Labels }}
<ul>
  {{ range $i, $label := .Labels.Nodes }}
  <li><a href="label/{{ $label.Name }}">{{ $label.Name }}
      {{ if $label.Discussions }}
      ({{ $label.Discussions.TotalCount }})
      {{ else }}
      (0)
      {{ end }}
    </a></li>
  {{ end }}
</ul>
{{ else }}
<h3>没有标签</h3>
{{ end }}
{{end}}