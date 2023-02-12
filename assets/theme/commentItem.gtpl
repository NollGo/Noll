{{define "CommentItemTemplate"}}
<div style="display: flex;">
  <img src="{{ .Author.AvatarURL }}" style="width: 3rem; height: 3rem;"/>
  <div class="flex-fill">
    <a href="{{ .Author.GitHubURL }}"> {{ .Author.ShowName }} </a>
    {{ .BodyHTML }}
  </div>
</div>
{{end}}