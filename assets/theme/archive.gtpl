<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate"}}
  <title>归档 —— {{ .Viewer.Name }}'s Blog </title>
</head>

<body>
  {{ if .Data }}
  {{ template "HeaderTemplate" . }}
  <div>
    <h1>归档</h1>
    {{ template "DiscussionGroupTemplate" .Data }}
    <ul class="ul" style="margin-left: -10px;">
      {{ if .Data.PageInfo.HasPrevPage }}
      <li class="li">
        <a href="/archive/{{ .Data.PageInfo.StartCursor }}.html">上一页</a>
      </li>
      {{ end }}
      {{ if .Data.PageInfo.HasNextPage }}
      <li class="li" style="float: right;">
        <a href="/archive/{{ .Data.PageInfo.EndCursor }}.html">下一页</a>
      </li>
      {{ end }}
    </ul>
  </div>
  {{ end }}
  {{ template "footerTemplate" .Viewer }}
</body>

</html>