{{define "HeaderTemplate"}}
<header>
<ul class="ul" style="margin-left: -10px;">
  <li class="li"><a href='{{ url "/" }}'>
      <h4 style="display: contents;">🏠 首页</h4>
    </a></li>
  <li class="li"><a href='{{ url "Archive" }}'>
      <h4 style="display: contents;">🗂️ 归档</h4>
    </a></li>
  <li class="li"><a href='{{ url "Categories" }}'>
      <h4 style="display: contents;">📑 分类</h4>
    </a></li>
  <li class="li"><a href='{{ url "Labels" }}'>
      <h4 style="display: contents;">🏷️ 标签</h4>
    </a></li>
  <li class="li"><a href='{{ url "About" }}'>
      <h4 style="display: contents;">👉 About</h4>
    </a></li>
</ul>
</header>
{{end}}