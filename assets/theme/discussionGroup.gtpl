{{define "DiscussionGroupTemplate"}}
<ul class="ul" style="margin-left: -10px;">
  {{ range $i, $discussion := .Nodes }}
  <li style="display: block; margin: 5px 0;"><a href="post/{{ $discussion.Number }}.html">{{ $discussion.Title }}
    </a>
    <ul style="display: contents;">
      <li class="li"><span class="COMMENT">
          {{ $discussion.Comments.TotalCount }}</span></li>
      {{ if $discussion.ReactionGroups }}
      {{ range $reaction := $discussion.ReactionGroups }}
      {{ if $reaction.Reactors.TotalCount }}
      <li class="li"><span class="{{ $reaction.Content }}">
          {{ $reaction.Reactors.TotalCount }}</span>
      </li>
      {{ end }}
      {{ end }}
      {{ end }}
    </ul>
  </li>
  {{ end }}
</ul>
{{end}}