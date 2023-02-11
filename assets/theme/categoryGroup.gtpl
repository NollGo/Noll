{{define "CategoryGroupTemplate"}}
{{ if .Categories }}
<ul class="ul" style="margin-left: -10px;">
  {{ range $i, $category := .Categories.Nodes }}
  {{ if $category.Discussions.TotalCount }}
  <li class="li"><a href="category/{{ $category.Name }}.html">
      {{ $category.EmojiHTML }} {{ $category.Name }}
    </a></li>
  {{ end }}
  {{ end }}
</ul>
{{ end }}
{{end}}