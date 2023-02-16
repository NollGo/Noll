<script>
  var alist = document.getElementsByTagName('a')
  alist.forEach(a => {
    a.href = fullpath(a.href)
  });

  function fullpath(path) {
    return '{{ .BaseURL }}' + path
  }
</script>