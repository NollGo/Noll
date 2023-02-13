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
    <div id="container" style="width: 100%; height: 500px; position: relative;"></div>
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

  <script async src="https://unpkg.com/es-module-shims@1.3.6/dist/es-module-shims.js"></script>

  <script type="importmap">
    {
      "imports": {
        "three": "https://unpkg.com/three@0.149.0/build/three.module.js",
        "threeAddons/": "https://unpkg.com/three@0.149.0/examples/jsm/"
      }
    }
  </script>

  <script type="module">

    import * as THREE from 'three';

    import Stats from 'threeAddons/libs/stats.module.js';

    import { STLLoader } from 'threeAddons/loaders/STLLoader.js';

    import { OrbitControls } from 'threeAddons/controls/OrbitControls.js';

    var string = `solid cube_corner
facet normal 0.0 -1.0 0.0
  outer loop
    vertex 0.0 0.0 0.0
    vertex 1.0 0.0 0.0
    vertex 0.0 0.0 1.0
  endloop
endfacet
facet normal 0.0 0.0 -1.0
  outer loop
    vertex 0.0 0.0 0.0
    vertex 0.0 1.0 0.0
    vertex 1.0 0.0 0.0
  endloop
endfacet
facet normal -1.0 0.0 0.0
  outer loop
    vertex 0.0 0.0 0.0
    vertex 0.0 0.0 1.0
    vertex 0.0 1.0 0.0
  endloop
endfacet
facet normal 0.577 0.577 0.577
  outer loop
    vertex 1.0 0.0 0.0
    vertex 0.0 1.0 0.0
    vertex 0.0 0.0 1.0
  endloop
endfacet
endsolid`;

    let container, stats;

    let camera, cameraTarget, scene, renderer;

    init();
    animate();

    function init() {

      container = document.getElementById('container')
      document.body.appendChild(container);

      camera = new THREE.PerspectiveCamera(35, window.innerWidth / window.innerHeight, 1, 15);
      camera.position.set(3, 0.15, 3);

      cameraTarget = new THREE.Vector3(0, - 0.25, 0);

      scene = new THREE.Scene();
      scene.background = new THREE.Color(0x72645b);
      scene.fog = new THREE.Fog(0x72645b, 2, 15);

      // Ground

      const plane = new THREE.Mesh(
        new THREE.PlaneGeometry(40, 40),
        new THREE.MeshPhongMaterial({ color: 0x999999, specular: 0x101010 })
      );
      plane.rotation.x = - Math.PI / 2;
      plane.position.y = - 0.5;
      scene.add(plane);

      // plane.receiveShadow = true;

      // ASCII string
      const loader = new STLLoader();
      var buffer = new TextEncoder().encode(string).buffer;
      var result = loader.parse(buffer)
      console.log(result);

      const material = new THREE.MeshPhongMaterial({ color: 0xff5533, specular: 0x111111, shininess: 200 });
      const mesh = new THREE.Mesh(result, material);

      mesh.position.set(0, - 0.25, 0.6);
      mesh.rotation.set(0, - Math.PI / 2, 0);
      mesh.scale.set(0.5, 0.5, 0.5);

      mesh.castShadow = true;
      mesh.receiveShadow = true;

      scene.add(mesh);

      // Lights

      scene.add(new THREE.HemisphereLight(0xffffff, 0xffffff));

      // addShadowedLight(1, 1, 1, 0xffffff, 1.35);
      // addShadowedLight(0.5, 1, - 1, 0xffaa00, 1);
      // renderer

      renderer = new THREE.WebGLRenderer({ antialias: true });
      renderer.setPixelRatio(window.devicePixelRatio);
      renderer.setSize(container.offsetWidth, container.offsetHeight);
      renderer.outputEncoding = THREE.sRGBEncoding;

      renderer.shadowMap.enabled = true;

      var controls = new OrbitControls(camera, renderer.domElement);
      controls.addEventListener('change', render);

      container.appendChild(renderer.domElement);

      // stats
      stats = new Stats();
      stats.dom.style = 'position: absolute; top: 0; left: 0;'
      container.appendChild(stats.dom);

      // window resize

      window.addEventListener('resize', onWindowResize);

    }

    function addShadowedLight(x, y, z, color, intensity) {

      const directionalLight = new THREE.DirectionalLight(color, intensity);
      directionalLight.position.set(x, y, z);
      scene.add(directionalLight);

      directionalLight.castShadow = true;

      const d = 1;
      directionalLight.shadow.camera.left = - d;
      directionalLight.shadow.camera.right = d;
      directionalLight.shadow.camera.top = d;
      directionalLight.shadow.camera.bottom = - d;

      directionalLight.shadow.camera.near = 1;
      directionalLight.shadow.camera.far = 4;

      directionalLight.shadow.bias = - 0.002;

    }

    function onWindowResize() {
      camera.aspect = window.innerWidth / window.innerHeight;
      camera.updateProjectionMatrix();

      renderer.setSize(container.offsetWidth, container.offsetHeight);
    }

    function animate() {
      requestAnimationFrame(animate);

      render();
      stats.update();
    }

    function render() {
      const timer = Date.now() * 0.0005;

      camera.position.x = Math.cos(timer) * 3;
      camera.position.z = Math.sin(timer) * 3;

      camera.lookAt(cameraTarget);

      renderer.render(scene, camera);

    }
  </script>
</body>

</html>