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
  </div>
  {{ end }}
</body>

</html>