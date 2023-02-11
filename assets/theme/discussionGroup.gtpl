{{define "DiscussionGroupTemplate"}}
<ul class="ul" style="margin-left: -10px;">
  {{ $time := 0 }}
  {{ range $i, $discussion := .Nodes }}
  {{ if ne $time $discussion.UpdatedAt.Year }}
  <a><time style="color: #7c7c7c;">
      <h3 style="margin: 0 0 0.3em 0;">{{ .UpdatedAt.Month }}/{{ .UpdatedAt.Year }}</h3>
    </time></a>
  {{ $time = $discussion.UpdatedAt.Year }}
  {{ end }}
  {{ template "DiscussionItemTemplate" $discussion }}
  {{ end }}
</ul>
{{end}}