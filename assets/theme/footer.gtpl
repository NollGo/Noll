{{define "footerTemplate"}}
<footer>
  {{ template "TopComponentTemplate" }}
  <div><a href='{{ url "NewPost" }}'>Create a post</a>
    | Powered by <a href="https://github.com/NollGo/Noll">Noll</a>
    | Theme default</div>
  <div><a href='{{ url "/" }}'>{{ .ShowName }}</a>© 2023 |
    <a href="https://creativecommons.org/licenses/by-nc/4.0/">CC BY-NC 4.0</a>
    | <a href='{{ url "RSS" }}'>RSS</a>
  </div>
</footer>
{{end}}