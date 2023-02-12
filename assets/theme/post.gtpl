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
    <div style="display: flex;">
      <label>{{ .Data.Category.Name }}</label>
    </div>
    {{ if .Data.Labels }}
    <div>
      {{ range $label := .Data.Labels.Nodes }}
      {{ template "LabelItemTemplate" $label }}
      {{ end }}
    </div>
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