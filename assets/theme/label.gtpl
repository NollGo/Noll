<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title>{{ .Data.Name }} Label —— {{ .Viewer.ShowName }}'s Blog </title>
</head>

<body {{ template "ColorStyleTemplate" .Data.Color }}>
  {{ template "HeaderTemplate" . }}
  <div>
    <h1> #{{ .Data.Name }} </h1>
    {{ .Data.Description }}
    {{ template "DiscussionGroupTemplate" .Data.Discussions }}
  </div>
  {{ template "footerTemplate" .Viewer }}
</body>

</html>