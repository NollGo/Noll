{{define "DiscussionGroup2Template"}}
<!-- 最多显示 7 条数据，并显示“更多文章”链接 -->
<ul class="ul" style="margin-left: -10px;">
  {{ range $i, $discussion := .Nodes }}
  {{ if lt $i 7 }}
  {{ template "DiscussionItemTemplate" $discussion }}
  {{ end }}
  {{ end }}
</ul>
<ul class="ul" style="margin-left: -10px;">
  <li class="li">
    <a href="/archive/1.html">更多文章 >>></a>
  </li>
</ul>
{{end}}