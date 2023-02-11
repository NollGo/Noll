{{define "DiscussionGroupTemplate"}}
<ul class="ul" style="margin-left: -10px;">
  {{ $time := time }}
  {{ range $i, $discussion := .Nodes }}
  {{ if ism $time $discussion.CreatedAt }}
  {{ else }}
  <a><time style="color: #c6c6c6;">
      <h3 style="margin: 0.5em 0 0.3em 0;">{{ $discussion.CreatedAt.Month }} {{ $discussion.CreatedAt.Year }}</h3>
    </time></a>
  {{ $time = $discussion.CreatedAt }}
  {{ end }}
  {{ template "DiscussionItemTemplate" $discussion }}
  {{ end }}
</ul>
{{end}}