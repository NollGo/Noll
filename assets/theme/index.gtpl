<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeaderTemplate"}}
  <title>{{ .Viewer.Name }}'s Blog </title>
</head>

<body>
  {{ if .Data }}
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