<!DOCTYPE html>
<html lang="zh-CN">

{{ $githubURL := .Data.GitHubURL }}

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title> {{ .Data.Title }}</title>
  <link rel="stylesheet" href="https://sindresorhus.com/github-markdown-css/github-markdown.css">
  <style>
    .mermaid {
      text-align: center;
      background-color: transparent !important;
    }

    article:first-of-type {
      margin-top: 40px;
    }

    table {
      width: 100% !important;
      min-width: 100% !important;
      display: table !important;
    }

    .markdown-body a {
      border-radius: 0;
      padding: 0;
      display: inline-block;
    }

    .markdown-body a:hover,
    .markdown-body a:active {
      background-color: transparent;
    }

    .reaction+.reaction {
      margin-left: 0;
    }

    .reaction a {
      border-radius: 100px;
    }

    .comment {
      width: 100%;
    }

    .comment-input {
      text-align: center;
      border: 1px solid #ddd;
      background-color: #f9f9f9;
      min-width: 100%;
      padding: 30px 0;
    }
  </style>
</head>

<body>
  {{ template "HeaderTemplate" . }}
  <div class="clearfix">
    <h1 style="margin-bottom: 0.5rem;"> {{ .Data.Title }} </h1>
    <div style="font-size: 1rem; display: flex; align-items: center;">
      <img src="{{ .Viewer.AvatarURL }}" style="width: 1.4rem; height: 1.4rem;" />
      <a href="/">{{ .Viewer.ShowName }}</a>
      发布在<a href="/category/{{ .Data.Category.Name }}.html">{{ .Data.Category.Name }}</a>
      于<time style="margin-left: 5px" title="{{ .Data.CreatedAt }}">
        {{ .Data.CreatedAt.Format "01-02-2006" }}</time>
    </div>
    <div id="map" style="height: 500px;"></div>
  </div>
  <article class="markdown-body" style="font-size: 1.2rem;">
    {{ .Data.BodyHTML }}
  </article>
  <ul class="ul" style="margin: 30px -10px;">
    <li class="li">{{ template "CategoryItemTemplate" .Data.Category }}</li>
    {{ if .Data.Labels }}
    {{ range $i, $label := .Data.Labels.Nodes }}
    <li class="li">{{ template "LabelItemTemplate" $label }}</li>
    {{ end }}
    {{ end }}
  </ul>
  <ul class="ul" style="text-align: center; margin: 30px auto;">
    {{ if .Data.UpvoteCount }}
    <li class="li reaction"><a href="{{ $githubURL }}">
        <span>{{ template "VoteSVGTemplate" 26 }} {{ .Data.UpvoteCount }}</span></a>
    </li>
    {{ end }}
    <li class="li reaction">
      <a href="{{ $githubURL }}"><span class="SMILING"></span></a>
    </li>
    {{ range $reaction := .Data.ReactionGroups }}
    {{ if $reaction.Reactors.TotalCount }}
    <li class="li reaction">
      <a href="{{ $githubURL }}"><span class="{{ $reaction.Content }}">
          {{ $reaction.Reactors.TotalCount }}</span></a>
    </li>
    {{ end }}
    {{ end }}
  </ul>
  <div style="display: flex; align-items: center; margin: 30px auto;">
    <div style="flex: 1; height: 1px; background-color: #ddd;"></div>
    <span class="COMMENT" style="margin: 0 12px"></span>
    <div style="flex: 1; height: 1px; background-color: #ddd;"></div>
  </div>
  {{ if .Data.Comments }}
  <ul class="ul" style="margin: 30px auto; font-size: 1rem;">
    {{ range $comment := .Data.Comments.Nodes }}
    <li class="li comment">
      {{ template "CommentItemTemplate" $comment }}
    </li>
    {{ end }}
  </ul>
  {{ end }}
  <a href="{{ $githubURL }}#reply" class="comment-input">前往 GitHub Discussion 评论</a>
  {{ template "footerTemplate" .Viewer }}
  <script>
    var mermaidEls = document.getElementsByClassName('notranslate')
    for (let index = 0; index < mermaidEls.length; index++) {
      const element = mermaidEls[index];
      if (element.getAttribute('lang') === 'mermaid') {
        element.classList.add('mermaid')
      }
    }
    var loaderEls = document.getElementsByClassName('js-render-enrichment-loader')
    for (let index = 0; index < loaderEls.length;) {
      const element = loaderEls[index];
      // removeChild 后，loaderEls 中也失去该 child
      element.parentElement.removeChild(element)
    }
  </script>
  <script src="https://cdn.jsdelivr.net/npm/mermaid@9/dist/mermaid.min.js"></script>
  <script>
    // https://docs.mathjax.org/en/latest/web/start.html
    MathJax = {
      tex: {
        inlineMath: [['$', '$'], ['\\(', '\\)']]
      }
    };
  </script>
  <script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-chtml.js"></script>
  <link rel="stylesheet" href="https://unpkg.com/leaflet@1.3.1/dist/leaflet.css"
    integrity="sha512-Rksm5RenBEKSKFjgI3a41vrjkw4EVPlJ3+OiI65vTjIdo9brlAacEuKOiQ5OFh7cOI1bkDwLqdLw3Zg0cRJAAQ=="
    crossorigin="" />
  <script src="https://unpkg.com/leaflet@1.3.1/dist/leaflet.js"
    integrity="sha512-/Nsx9X4HebavoBvEBuyp3I7od5tA0UzAxs+j83KgC8PU0kgB4XiK4Lfe4y4cgBtaRJQEIFCW+oC506aPT2L1zw=="
    crossorigin=""></script>
  <script src="https://unpkg.com/topojson@3.0.2/dist/topojson.min.js"></script>
  <script>
    function maprender(target, dataType, data) {
      // 初始化地图
      var map = L.map(target).setView([48.505, -0.09], 7);

      // 加载地图底图
      // https://cartodb-basemaps-{s}.global.ssl.fastly.net/light_all/{z}/{x}/{y}.png
      let bgLayerPositron = L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a> &copy; <a href="http://cartodb.com/attributions">CartoDB</a>',
        subdomains: 'abcd',
        maxZoom: 19
      });

      bgLayerPositron.addTo(map);

      //extend Leaflet to create a GeoJSON layer from a TopoJSON file
      L.TopoJSON = L.GeoJSON.extend({
        addData: function (data) {
          var geojson, key;
          if (data.type === "Topology") {
            for (key in data.objects) {
              if (data.objects.hasOwnProperty(key)) {
                geojson = topojson.feature(data, data.objects[key]);
                L.GeoJSON.prototype.addData.call(this, geojson);
              }
            }
            return this;
          }
          L.GeoJSON.prototype.addData.call(this, data);
          return this;
        }
      });
      L.topoJson = function (data, options) {
        return new L.TopoJSON(data, options);
      };

      if (dataType === 'geojson') {
        // 加载 GeoJSON 数据
        var geojson = L.geoJson(null, {
          style: function (feature) {
            let color = feature.properties.color | '#35495d'
            return { color: color };
          }
        }).addTo(map)
        let area = geojson.addData(data);
        // 定位到当前加载数据的区域
        map.fitBounds(area.getBounds());
      } else {
        //create an empty geojson layer
        //with a style and a popup on click
        var geojson = L.topoJson(null, {
          style: function (feature) {
            return {
              color: "#000",
              opacity: 1,
              weight: 1,
              fillColor: '#35495d',
              fillOpacity: 0.8
            }
          },
          onEachFeature: function (feature, layer) {
            if (feature.properties.name) {
              layer.bindPopup('<p>' + feature.properties.name + '</p>')
            }
          }
        }).addTo(map);
        let area = geojson.addData(data);
        // 定位到当前加载数据的区域
        map.fitBounds(area.getBounds());
      }
    }
    var maps = document.getElementsByTagName('section')
    for (let index = 0; index < maps.length; index++) {
      const element = maps[index]
      let dataType = element.getAttribute('data-type')
      if (dataType === 'geojson' || dataType === 'topojson') {
        let target = element.firstElementChild
        let data = JSON.parse(target.getAttribute('data-plain'))
        target.style.height = '400px'
        target.innerHTML = ''
        maprender(target, dataType, data)
      }
    }
  </script>
</body>

</html>