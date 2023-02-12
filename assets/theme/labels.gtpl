<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title>标签 —— {{ .Viewer.Name }}'s Blog </title>
</head>

<body>
  <style>
    .flex-fill:not(:last-of-type) {
      margin-right: 20px;
    }
  </style>
  {{ template "HeaderTemplate" . }}
  {{ if .Labels.TotalCount }}
  <div class="column">
    {{ range $label := .Labels.Nodes }}
    {{ if $label.Discussions.TotalCount }}
    <div class="flex-fill" {{ template "ColorStyleTemplate" $label.Color }}>
      <h1>#{{ $label.Name }} ({{ $label.Discussions.TotalCount }})</h1>
      {{ template "DiscussionGroup2Template" $label.Discussions }}
      <ul class="ul" style="margin-left: -10px;">
        <li class="li">
          <a href="/label/{{ $label.Name }}.html">更多文章 >>></a>
        </li>
      </ul>
    </div>
    {{ end }}
    {{ end }}
  </div>
  {{ else }}
  <h4>这里还没有标签</h4>
  {{ end }}
  {{ template "footerTemplate" .Viewer }}
</body>

</html>