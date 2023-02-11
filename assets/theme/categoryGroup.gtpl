{{define "CategoryGroupTemplate"}}
{{ if .Categories }}
<ul class="ul">
  {{ range $i, $category := .Categories.Nodes }}
  <li class="li"><a href="category/{{ $category.Name }}">
      {{ $category.EmojiHTML }} {{ $category.Name }}
    </a></li>
  {{ end }}
</ul>
{{ end }}
{{end}}