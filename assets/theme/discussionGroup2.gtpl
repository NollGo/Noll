{{define "DiscussionGroup2Template"}}
<!-- 最多显示 7 条数据 -->
<ul class="ul" style="margin-left: -10px;">
  {{ range $i, $discussion := .Nodes }}
  {{ if lt $i 7 }}
  {{ template "DiscussionItemTemplate" $discussion }}
  {{ end }}
  {{ end }}
</ul>
{{end}}