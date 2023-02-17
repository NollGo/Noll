<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title>归档 —— {{ .Viewer.ShowName }}'s Blog </title>
</head>

<body>
  {{ if .Data }}
  {{ template "HeaderTemplate" . }}
  <div class="clearfix">
    <h1>归档</h1>
    {{ template "DiscussionGroupTemplate" .Data }}
    <ul class="ul" style="margin-left: -10px;">
      {{ if .Data.PageInfo.HasPrevPage }}
      <li class="li">
        <a href='{{ url2 .Data .Data.PageInfo.StartCursor }}'>上一页</a>
      </li>
      {{ end }}
      {{ if .Data.PageInfo.HasNextPage }}
      <li class="li" style="float: right;">
        <a href='{{ url2 .Data .Data.PageInfo.EndCursor }}'>下一页</a>
      </li>
      {{ end }}
    </ul>
  </div>
  {{ end }}
  {{ template "footerTemplate" .Viewer }}
</body>

</html>