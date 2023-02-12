<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title>分类 —— {{ .Viewer.Name }}'s Blog </title>
</head>

<body>
  {{ template "HeaderTemplate" . }}
  {{ if .Categories.TotalCount }}
  <div class="column">
    {{ range $category := .Categories.Nodes }}
    {{ if $category.Discussions.TotalCount }}
    <div class="flex-fill">
      <h1>{{ $category.EmojiHTML }} {{ $category.Name }} ({{ $category.Discussions.TotalCount }})</h1>
      {{ template "DiscussionGroup2Template" $category.Discussions }}
      <ul class="ul" style="margin-left: -10px;">
        <li class="li">
          <a href="/category/{{ $category.Name }}.html">更多文章 >>></a>
        </li>
      </ul>
    </div>
    {{ end }}
    {{ end }}
  </div>
  {{ else }}
  <h4>这里还没有分类</h4>
  {{ end }}
  {{ template "footerTemplate" .Viewer }}
</body>

</html>