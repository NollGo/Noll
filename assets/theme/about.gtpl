<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title>Aoubt —— {{ .Viewer.Name }}'s Blog </title>
</head>

<body>
  {{ template "HeaderTemplate" . }}
  <h1>About {{ .Viewer.Name }}</h1>
  <p>{{ .Viewer.Bio }}</p>
  <p>🏢 {{ .Viewer.Company }}</p>
  <p>🌍 {{ .Viewer.Location }}</p>
  <p>📧 {{ .Viewer.Email }}</p>
  <p>😺 <a style="padding: 0px;" href="{{ .Viewer.GitHubURL }}">{{ .Viewer.GitHubURL }}</a></p>
  <!-- style="white-space: pre-wrap;" -->
  {{ template "footerTemplate" .Viewer }}
</body>

</html>