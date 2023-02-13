{{define "CommentItemTemplate"}}
<div style="display: flex;">
  <div>
    <a class="text" href="{{ .Author.GitHubURL }}">
      <img src="{{ .Author.AvatarURL }}" style="width: 3rem; height: 3rem;" /></a>
  </div>
  <div style="margin-left: 10px; flex: 1;">
    <div style="display: flex; color: #555;">
      <a class="text" href="{{ .Author.GitHubURL }}">{{ .Author.ShowName }}</a>
      <time style="margin: 0 10px" title="{{ .UpdatedAt }}"> {{ .UpdatedAt.Format "01-02" }} </time>
      <div style="flex: 1;"></div>
      <a class="text" href="{{ .GitHubURL }}">回复</a>
    </div>
    <div class="markdown-body" style="margin: 12px 0;">{{ .BodyHTML }}</div>
    <ul class="ul column">
      <li class="li">
        <span>{{ template "VoteSVGTemplate" 22 }} {{ .UpvoteCount }}</span>
      </li>
      {{ range $reaction := .ReactionGroups }}
      {{ if $reaction.Reactors.TotalCount }}
      <li class="li" style="margin-left: 10px">
        <span class="{{ $reaction.Content }}"> {{ $reaction.Reactors.TotalCount }}</span>
      </li>
      {{ end }}
      {{ end }}
    </ul>
  </div>
</div>
{{end}}