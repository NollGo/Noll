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
      {{ template "DiscussionGroupTemplate" .Data }}
      <div>
        <div style="display: block; margin: 1em -10px;">
          <a href="/archive.html">更多文章……</a>
        </div>
      </div>
    </div>
    <div>
      <h1>标签 ({{ .Labels.TotalCount }})</h1>
      {{ template "LabelGroupTemplate" .Labels }}
    </div>
  </div>
  {{ end }}
</body>

</html>