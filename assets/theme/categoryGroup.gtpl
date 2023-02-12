{{define "CategoryGroupTemplate"}}
{{ if . }}
<ul class="ul" style="margin-left: -10px;">
  {{ range $i, $category := .Nodes }}
  {{ if $category.Discussions.TotalCount }}
  <li class="li">{{ template "CategoryItemTemplate" $category }}</li>
  {{ end }}
  {{ end }}
</ul>
{{ end }}
{{end}}