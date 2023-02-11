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
      <h1>近期文章</h1>
      {{ template "DiscussionGroup2Template" .Data }}
    </div>
    <div>
      <h1>标签 ({{ .Labels.TotalCount }})</h1>
      {{ template "LabelGroupTemplate" .Labels }}
    </div>
  </div>
  {{ end }}
  {{ template "footerTemplate" .Viewer }}
</body>

</html>