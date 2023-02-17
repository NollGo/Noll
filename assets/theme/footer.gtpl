{{define "footerTemplate"}}
<footer>
  <div>Powered by <a href="https://github.com/ThreeTenth/GitHub-Discussions-to-Blog">toPages</a> | Theme default</div>
  <div><a href='{{ url "/" }}'>{{ .Name }}</a>© 2023 |
    <a href="https://creativecommons.org/licenses/by-nc/4.0/">CC BY-NC 4.0</a>
    <!-- | <a href='{{ url "RSS" }}'>RSS</a> -->
  </div>
</footer>
{{end}}