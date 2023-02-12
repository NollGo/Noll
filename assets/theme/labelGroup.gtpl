{{define "LabelGroupTemplate"}}
{{ if . }}
<ul class="ul" style="margin-left: -10px;">
  {{ range $i, $label := .Nodes }}
  {{ if $label.Discussions.TotalCount }}
  <li class="li">{{ template "LabelItemTemplate" $label }}</li>
  {{ end }}
  {{ end }}
</ul>
{{ end }}
{{end}}