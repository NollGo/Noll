package main

import (
	"fmt"
	"github.com/howeyc/fsnotify"
	"github.com/lxzan/gws"
	"net/http"
	"os"
	"path/filepath"
)

var upgrader = func(event gws.Event) *gws.Upgrader {
	return gws.NewUpgrader(func(c *gws.Upgrader) {
		c.CompressEnabled = true
		c.CheckTextEncoding = true
		c.MaxContentLength = 32 * 1024 * 1024
		c.EventHandler = event
	})
}

func debugWs(config Config, _render func() error) http.Handler {
	websocket := &DebugWs{}

	// watch file change config.ThemeDir
	dirList := collDir(config.ThemeDir)
	for i := range dirList {
		watch(dirList[i], _render, websocket)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		socket, err := upgrader(websocket).Accept(w, r)
		if err != nil {
			return
		}

		websocket.socket = socket
		go socket.Listen()
	})
}

func watch(dir string, _render func() error, websocket *DebugWs) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	if err := watcher.Watch(dir); err == nil {
		go func() {
			fmt.Println("Start watch file change", dir)
			for {
				select {
				case <-watcher.Event:
					if err := _render(); err != nil {
						panic(err)
					}

					if websocket.socket != nil {
						_ = websocket.socket.WriteString("reload")
					}
				case err := <-watcher.Error:
					fmt.Println("error:", err)
				}
			}
		}()
	}
}

// collect all dir
func collDir(path string) []string {
	if path == "" {
		return []string{}
	}

	var dirs []string

	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirs = append(dirs, path)
		}
		return nil
	}); err != nil {
		return []string{}
	}

	return dirs
}

type DebugWs struct {
	socket *gws.Conn
}

func (d DebugWs) OnOpen(socket *gws.Conn) {
}

func (d DebugWs) OnError(socket *gws.Conn, err error) {
}

func (d DebugWs) OnClose(socket *gws.Conn, code uint16, reason []byte) {
}

func (d DebugWs) OnPing(socket *gws.Conn, payload []byte) {
}

func (d DebugWs) OnPong(socket *gws.Conn, payload []byte) {
}

func (d DebugWs) OnMessage(socket *gws.Conn, message *gws.Message) {
}
