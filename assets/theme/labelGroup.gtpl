{{define "LabelGroupTemplate"}}
<h1>标签({{ .Labels.TotalCount }})</h1>
{{ if .Labels }}
<ul class="ul" style="margin-left: -10px;">
  {{ range $i, $label := .Labels.Nodes }}
  <li class="li"><a href="label/{{ $label.Name }}.html">
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