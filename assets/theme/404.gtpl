<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title>404 Not Found —— {{ .Viewer.ShowName }}'s Blog </title>
</head>

<body>
  {{ template "HeaderTemplate" . }}
  <div class="clearfix"></div>
  <h4>404 Not Found</h4>
  <div><a href='#' onclick="javascript :history.back(-1); return false;">返回上页</a></div>
  <div><a href='{{ url " /" }}'>返回首页</a></div>
  {{ template "footerTemplate" .Viewer }}
</body>

</html>