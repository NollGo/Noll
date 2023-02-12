<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title>{{ .Viewer.ShowName }}'s Blog </title>
</head>

<body>
  {{ if .Data }}
  {{ template "HeaderTemplate" . }}
  <div class="column">
    <div class="flex-fill">
      <h1>近期文章</h1>
      {{ template "DiscussionGroup2Template" .Data }}
      <ul class="ul" style="margin-left: -10px;">
        <li class="li">
          <a href="/archive/1.html">更多文章 >>></a>
        </li>
      </ul>
    </div>
    <div>
      <h1>分类</h1>
      {{ template "CategoryGroupTemplate" .Categories }}
    </div>
    <div>
      <h1>标签</h1>
      {{ template "LabelGroupTemplate" .Labels }}
    </div>
  </div>
  {{ end }}
  {{ template "footerTemplate" .Viewer }}
</body>

</html>