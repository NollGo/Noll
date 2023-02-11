{{define "LabelGroupTemplate"}}
{{ if . }}
<ul class="ul" style="margin-left: -10px;">
  {{ range $i, $label := .Nodes }}
  {{ if $label.Discussions.TotalCount }}
  <li class="li"><a href="/label/{{ $label.Name }}.html">
      {{ template "LabelTemplate" $label }}
      ({{ $label.Discussions.TotalCount }})
    </a></li>
  {{ end }}
  {{ end }}
</ul>
{{ end }}
{{end}}