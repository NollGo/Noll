<!DOCTYPE html>
<html lang="zh-CN">

{{ $githubURL := .Data.GitHubURL }}

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title> {{ .Data.Title }}</title>
  <style>
    .markdown p {
      line-height: 1.4;
    }

    .reaction+.reaction {
      margin-left: 0;
    }

    .reaction a {
      border-radius: 100px;
    }
  </style>
</head>

<body>
  {{ template "HeaderTemplate" . }}
  <div class="clearfix">
    <h1 style="margin-bottom: 0.5rem;"> {{ .Data.Title }} </h1>
    <div style="font-size: 1rem; display: flex; align-items: center;">
      <img src="{{ .Viewer.AvatarURL }}" style="width: 1.4rem; height: 1.4rem;" />
      <a href="/">{{ .Viewer.ShowName }}</a>
      发布在<a href="/category/{{ .Data.Category.Name }}.html">{{ .Data.Category.Name }}</a>
      于<time style="margin-left: 5px" title="{{ .Data.CreatedAt }}">
        {{ .Data.CreatedAt.Format "01-02-2006" }}</time>
    </div>
  </div>
  <article class="markdown">
    <div> {{ .Data.BodyHTML }} </div>
    {{ if .Data.ReactionGroups }}
    <ul class="ul" style="text-align: center; margin: 30px auto;">
      <li class="li reaction">
        <a href="{{ $githubURL }}"><span class="SMILING"></span></a>
      </li>
      {{ range $reaction := .Data.ReactionGroups }}
      {{ if $reaction.Reactors.TotalCount }}
      <li class="li reaction">
        <a href="{{ $githubURL }}"><span class="{{ $reaction.Content }}">
            {{ $reaction.Reactors.TotalCount }}</span></a>
      </li>
      {{ end }}
      {{ end }}
    </ul>
    {{ else }}
    {{ end }}
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
        {{ template "CommentItemTemplate" $comment }}
      </li>
      {{ end }}
      </div>
      {{ end }}
  </article>
  {{ template "footerTemplate" .Viewer }}
</body>

</html>