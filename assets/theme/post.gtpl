<!DOCTYPE html>
<html lang="zh-CN">

{{ $githubURL := .Data.GitHubURL }}

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title> {{ .Data.Title }}</title>
  <link rel="stylesheet" href="https://sindresorhus.com/github-markdown-css/github-markdown.css">
  <style>
    .mermaid {
      text-align: center;
      background-color: transparent !important;
    }

    article:first-of-type {
      margin-top: 40px;
    }

    table {
      width: 100% !important;
      min-width: 100% !important;
      display: table !important;
    }

    .markdown-body a {
      border-radius: 0;
      padding: 0;
      display: inline-block;
    }

    .markdown-body a:hover,
    .markdown-body a:active {
      background-color: transparent;
    }

    .reaction+.reaction {
      margin-left: 0;
    }

    .reaction a {
      border-radius: 100px;
    }

    .comment {
      width: 100%;
    }

    .comment-input {
      text-align: center;
      border: 1px solid #ddd;
      background-color: #f9f9f9;
      min-width: 100%;
      padding: 30px 0;
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
    <!-- <div id="container" style="width: 100%; height: 500px; position: relative;"></div> -->
  </div>
  <article class="markdown-body" style="font-size: 1.2rem;">
    {{ .Data.BodyHTML }}
  </article>
  <ul class="ul" style="margin: 30px -10px;">
    <li class="li">{{ template "CategoryItemTemplate" .Data.Category }}</li>
    {{ if .Data.Labels }}
    {{ range $i, $label := .Data.Labels.Nodes }}
    <li class="li">{{ template "LabelItemTemplate" $label }}</li>
    {{ end }}
    {{ end }}
  </ul>
  <ul class="ul" style="text-align: center; margin: 30px auto;">
    {{ if .Data.UpvoteCount }}
    <li class="li reaction"><a href="{{ $githubURL }}">
        <span>{{ template "VoteSVGTemplate" 26 }} {{ .Data.UpvoteCount }}</span></a>
    </li>
    {{ end }}
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
  <div style="display: flex; align-items: center; margin: 30px auto;">
    <div style="flex: 1; height: 1px; background-color: #ddd;"></div>
    <span class="COMMENT" style="margin: 0 12px"></span>
    <div style="flex: 1; height: 1px; background-color: #ddd;"></div>
  </div>
  {{ if .Data.Comments }}
  <ul class="ul" style="margin: 30px auto; font-size: 1rem;">
    {{ range $comment := .Data.Comments.Nodes }}
    <li class="li comment">
      {{ template "CommentItemTemplate" $comment }}
    </li>
    {{ end }}
  </ul>
  {{ end }}
  <a href="{{ $githubURL }}#reply" class="comment-input">前往 GitHub Discussion 评论</a>
  {{ template "footerTemplate" .Viewer }}
</body>

</html>