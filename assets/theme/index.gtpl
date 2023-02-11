<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate"}}
  <title>{{ .Viewer.Name }}'s Blog </title>
</head>

<body>
  {{ if .Data }}
  {{ template "HeaderTemplate" . }}
  {{ template "CategoryGroupTemplate" . }}
  <div class="column">
    <div class="flex-fill">
      {{ template "DiscussionGroupTemplate" . }}
    </div>
    <div>
      {{ template "LabelGroupTemplate" . }}
    </div>
  </div>
  {{ end }}
</body>

</html>