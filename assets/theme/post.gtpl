<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title> {{ .Data.Title }}</title>
</head>

<body>
  {{ template "HeaderTemplate" . }}
  <div class="clearfix"></div>
  <h1> {{ .Data.Title }} </h1>
  <time>{{ .Data.CreatedAt }}</time>
  <article class="markdown">
    <div> {{ .Data.BodyHTML }} </div>
    <ul class="ul" style="margin-left: -10px;">
      <li class="li">{{ template "CategoryItemTemplate" .Data.Category }}</li>
    </ul>
    {{ if .Data.Labels }}
    <ul class="ul" style="margin-left: -10px;">
      {{ range $i, $label := .Data.Labels.Nodes }}
      <li class="li">{{ template "LabelItemTemplate" $label }}</li>
      {{ end }}
    </ul>
    {{ end }}
    {{ if .Data.Comments }}
    <h4>去评论</h4>
    <div>
      {{ range $comment := .Data.Comments.Nodes }}
      <p>
        {{ $comment.Body }}
      </p>
      {{ end }}
    </div>
    {{ end }}
  </article>
  {{ template "footerTemplate" .Viewer }}
</body>

</html>