<style>
  .debug {
    display: flex;
    position: fixed;
    /* absolute */
    font-size: 1.0rem;
    padding: 4px 12px;
    border-radius: 0.2rem;
    background: #515151;
    box-shadow: 0px 0px 1px black;
    line-height: 1.4rem;
  }

  .debug .button {
    text-decoration: none;
    margin-left: 0.6rem;
    padding: 0.1rem;
  }

  .debug .button:hover {
    border-radius: 0.2rem;
    background: #252525;
  }

  .debug .drop::after {
    content: "::";
    color: #c6c6c6;
    font-weight: bold;
    cursor: grab;
  }

  .debug .loader {
    border: 2px solid #f3f3f3;
    border-top: 2px solid #555;
    border-radius: 50%;
    width: 1.2rem;
    height: 1.2rem;
    display: inline-block;
    animation: spin 2s linear infinite;
  }

  @keyframes spin {
    0% {
      transform: rotate(0deg);
    }

    100% {
      transform: rotate(360deg);
    }
  }
</style>
<script>
  let debugRefresh = createDebugButton('⚡', '/build', 'Refresh this page(F5)')
  let debugIncrease = createDebugButton('⛏️', '/build?mode=increase', 'Incremental update(F6)')
  let debugFull = createDebugButton('⚒️', '/build?mode=full', 'Full update(F7)')
  let isDebugMenuDrop = false
  window.addEventListener('DOMContentLoaded', function () {
    let debugMenu = document.createElement('div')
    debugMenu.style.right = '10px'
    debugMenu.style.top = '10px'
    debugMenu.classList.add('debug')

    let drop = document.createElement("span")
    drop.classList.add('drop')
    drop.onmousedown = function (e) {
      isDebugMenuDrop = true
      debugMenu.style.removeProperty('right')
      debugMenu.style.left = `${e.clientX - 16}px`
      debugMenu.style.top = `${e.clientY - 16}px`
    }
    window.onmouseup = function (e) {
      isDebugMenuDrop = false
    }
    window.onmousemove = function (e) {
      if (!isDebugMenuDrop) return
      debugMenu.style.left = `${e.clientX - 16}px`
      debugMenu.style.top = `${e.clientY - 16}px`
    }

    debugMenu.appendChild(drop)
    debugMenu.appendChild(debugRefresh)
    debugMenu.appendChild(debugIncrease)
    debugMenu.appendChild(debugFull)
    document.body.appendChild(debugMenu)

    window.addEventListener("keydown", (event) => {
      if ("F5" == event.key) {
        event.preventDefault()
        debugRefresh.click()
      } else if ("F6" == event.key) {
        event.preventDefault()
        debugIncrease.click()
      } else if ("F7" == event.key) {
        event.preventDefault()
        debugFull.click()
      }
    })
  })

  function createDebugButton(text, href, title) {
    let debugButton = document.createElement('a')
    debugButton.classList.add('button')
    debugButton.innerText = text
    debugButton.href = href
    debugButton.title = title
    let build = async function () {
      try {
        let resp = await fetch(href)
        if (resp.ok) {
          location.reload()
        } else {
          debugButton.classList.remove('loader')
          let body = await resp.text()
          throw new Error(`${resp.statusText}(${resp.status}), ${body}`)
        }
      } catch (e) {
        alert(e)
        debugButton.classList.remove('loader')
      }
    }
    debugButton.onclick = function () {
      debugButton.classList.add('loader')
      build()
      return false
    }
    return debugButton
  }

  function ws() {
    let webSocket = new WebSocket("ws://localhost:20000/ws");
    webSocket.onopen = function (event) {
      console.log("onopen", event);
    };

    webSocket.onmessage = function (event) {
      if (event.type === 'message' && event.data === 'reload') {
        console.log("reload")
        location.reload()
      }
    };
  }

  ws()
</script>