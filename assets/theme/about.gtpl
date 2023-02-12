<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title>Aoubt —— {{ .Viewer.ShowName }}'s Blog </title>
</head>

<body>
  {{ template "HeaderTemplate" . }}
  <div class="clearfix"></div>
  <h1>About {{ .Viewer.ShowName }}</h1>
  {{ if .Viewer.Bio }}
  <p>{{ .Viewer.Bio }}</p>
  {{ end }}
  {{ if .Viewer.Company }}
  <p>🏢 {{ .Viewer.Company }}</p>
  {{ end }}
  {{ if .Viewer.Location }}
  <p>🌍 {{ .Viewer.Location }}</p>
  {{ end }}
  {{ if .Viewer.Email }}
  <p>📧 {{ .Viewer.Email }}</p>
  {{ end }}
  <p>😺 <a style="padding: 0px;" href="{{ .Viewer.GitHubURL }}">{{ .Viewer.GitHubURL }}</a></p>
  {{ if .Viewer.Twitter }}
  <p>🕊️ <a style="padding: 0px;" href="https://twitter.com/{{ .Viewer.Twitter }}">
      https://twitter.com/{{ .Viewer.Twitter }}</a></p>
  {{ end }}
  <!-- style="white-space: pre-wrap;" -->
  {{ template "footerTemplate" .Viewer }}
</body>

</html>