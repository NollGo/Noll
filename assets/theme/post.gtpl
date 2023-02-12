<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title> {{ .Data.Title }}</title>
  <style>
    .markdown p {
      line-height: 1.4;
    }
  </style>
</head>

<body>
  {{ template "HeaderTemplate" . }}
  <div class="clearfix"></div>
  <h1> {{ .Data.Title }} </h1>
  <time title="{{ .Data.CreatedAt }}">{{ .Data.CreatedAt.Format "2006-01-02" }}</time>
  <article class="markdown">
    <div> {{ .Data.BodyHTML }} </div>
    <ul class="ul" style="margin-left: -10px;">
      <li class="li">{{ template "CategoryItemTemplate" .Data.Category }}</li>
      {{ if .Data.Labels }}
      {{ range $i, $label := .Data.Labels.Nodes }}
      <li class="li">{{ template "LabelItemTemplate" $label }}</li>
      {{ end }}
      {{ end }}
    </ul>
    {{ if .Data.Comments }}
    <h4>去评论</h4>
    <ul class="ul">
      {{ range $comment := .Data.Comments.Nodes }}
      <li class="li">
        <div>{{ $comment.BodyHTML }}</div>
      </li>
      {{ end }}
      </div>
      {{ end }}
  </article>
  {{ template "footerTemplate" .Viewer }}
</body>

</html>