{{define "CategoryGroupTemplate"}}
{{ if .Categories }}
<ul class="ul" style="margin-left: -10px;">
  {{ range $i, $category := .Categories.Nodes }}
  <li class="li"><a href="category/{{ $category.Name }}.html">
      {{ $category.EmojiHTML }} {{ $category.Name }}
    </a></li>
  {{ end }}
</ul>
{{ end }}
{{end}}