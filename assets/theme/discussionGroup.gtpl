{{define "DiscussionGroupTemplate"}}
<h1>文章({{ .Data.TotalCount }})</h1>
<ul>
  {{ range $i, $discussion := .Data.Nodes }}
  <li><a href="post/{{ $discussion.Number }}.html">{{ $discussion.Title }}
      ({{ $discussion.Comments.TotalCount }})
    </a></li>
  {{ if $discussion.ReactionGroups }}
  <ul style="list-style-type: none">
    {{ range $reaction := $discussion.ReactionGroups }}
    {{ if gt $reaction.Reactors.TotalCount 0 }}
    <li style="display: contents;"><span class="{{ $reaction.Content }}">
        {{ $reaction.Reactors.TotalCount }}</span>
    </li>
    {{ end }}
    {{ end }}
  </ul>
  {{ end }}
  {{ end }}
</ul>
{{end}}