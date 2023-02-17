<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title>标签 —— {{ .Viewer.ShowName }}'s Blog </title>
</head>

<body>
  {{ template "HeaderTemplate" . }}
  {{ if .Labels.TotalCount }}
  <div class="clearfix">
    {{ range $label := .Labels.Nodes }}
    {{ if $label.Discussions.TotalCount }}
    <div {{ template "ColorStyleTemplate" $label.Color }}>
      <h1>#{{ $label.Name }} ({{ $label.Discussions.TotalCount }})</h1>
      {{ template "DiscussionGroup2Template" $label.Discussions }}
      <ul class="ul" style="margin-left: -10px;">
        <li class="li">
          <a href="{{ url $label }}">更多文章 >>></a>
        </li>
      </ul>
    </div>
    {{ end }}
    {{ end }}
  </div>
  {{ else }}
  <h4>这里还没有标签</h4>
  {{ end }}
  {{ template "footerTemplate" .Viewer }}
</body>

</html>