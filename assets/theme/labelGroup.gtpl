{{define "LabelGroupTemplate"}}
<h1>标签({{ .Labels.TotalCount }})</h1>
{{ if .Labels }}
<ul class="ul">
  {{ range $i, $label := .Labels.Nodes }}
  <li class="li"><a href="label/{{ $label.Name }}">
      {{ template "LabelTemplate" $label }}
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