{{define "LabelGroupTemplate"}}
{{ if . }}
<ul class="ul" style="margin-left: -10px;">
  {{ range $i, $label := .Nodes }}
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
{{ end }}
{{end}}