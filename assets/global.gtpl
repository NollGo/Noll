<script>
  function fullpath(path) {
    return '{{ .BaseURL }}' + path
  }
</script>
{{ if .GamID }}
<!-- Google tag (gtag.js) -->
<script async src='https://www.googletagmanager.com/gtag/js?id={{ .GamID }}'></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag() { dataLayer.push(arguments); }
  gtag('js', new Date());

  gtag('config', '{{ .GamID }}');
</script>
{{ end }}