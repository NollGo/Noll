<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title>分类 —— {{ .Viewer.ShowName }}'s Blog </title>
</head>

<body>
  {{ template "HeaderTemplate" . }}
  {{ if .Categories.TotalCount }}
  <div class="clearfix">
    {{ range $category := .Categories.Nodes }}
    {{ if $category.Discussions.TotalCount }}
    <div>
      <h1>{{ $category.EmojiHTML }} {{ $category.Name }} ({{ $category.Discussions.TotalCount }})</h1>
      {{ template "DiscussionGroup2Template" $category.Discussions }}
      <ul class="ul" style="margin-left: -10px;">
        <li class="li">
          <a href="{{ url $category }}">更多文章 >>></a>
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