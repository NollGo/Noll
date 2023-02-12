{{define "DiscussionItemTemplate"}}
<li style="display: block; margin: 5px 0;">
  <a href="/post/{{ .Number }}.html">
    {{ .Title }}
    <ul style="display: contents;">
      <li class="li"><span class="COMMENT">
          {{ .Comments.TotalCount }}</span></li>
      {{ if .ReactionGroups }}
      {{ range $reaction := .ReactionGroups }}
      {{ if $reaction.Reactors.TotalCount }}
      <li class="li"><span class="{{ $reaction.Content }}">
          {{ $reaction.Reactors.TotalCount }}</span>
      </li>
      {{ end }}
      {{ end }}
      {{ end }}
    </ul>
  </a>
</li>
{{end}}