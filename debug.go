package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/lxzan/gws"
)

var upgrader = func(event gws.Event) *gws.Upgrader {
	return gws.NewUpgrader(event, &gws.ServerOption{
		CompressEnabled:     true,
		CheckUtf8Enabled:    true,
		ReadMaxPayloadSize:  32 * 1024 * 1024,
		WriteMaxPayloadSize: 32 * 1024 * 1024,
	})
}

func debugWs(config Config, _render func() error) http.Handler {
	websocket := &DebugWs{}

	// watch file change config.ThemeDir
	dirList := collDir(config.ThemeDir)
	pathChan, err := watch(_render, websocket)
	if err != nil {
		panic(err)
	}

	for i := range dirList {
		pathChan <- dirList[i]
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

func watch(_render func() error, websocket *DebugWs) (chan string, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:

				if event.Has(fsnotify.Write) {
					if err := _render(); err != nil {
						fmt.Println("error:", err)
					}

					if websocket.socket != nil {
						_ = websocket.socket.WriteString("reload")
					}
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()

	pathChan := make(chan string)
	go func() {
		for {
			select {
			case path := <-pathChan:
				if err := watcher.Add(path); err == nil {
					fmt.Println("Start watch file change", path)
				} else {
					panic(err)
				}
			}
		}
	}()

	return pathChan, nil
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

// DebugWs is 调试 websocket event
type DebugWs struct {
	socket *gws.Conn
}

// OnOpen is websocket 建立连接事件
func (d DebugWs) OnOpen(socket *gws.Conn) {
}

// OnError is websocket 错误事件
// IO错误, 协议错误, 压缩解压错误...
func (d DebugWs) OnError(socket *gws.Conn, err error) {
}

// OnClose is websocket 关闭事件
// 另一端发送了关闭帧
func (d DebugWs) OnClose(socket *gws.Conn, code uint16, reason []byte) {
}

// OnPing is websocket 心跳探测事件
func (d DebugWs) OnPing(socket *gws.Conn, payload []byte) {
}

// OnPong is websocket 心跳响应事件
func (d DebugWs) OnPong(socket *gws.Conn, payload []byte) {
}

// OnMessage is websocket 消息事件
// 如果开启了AsyncReadEnabled, 可以在一个连接里面并行处理多个请求
func (d DebugWs) OnMessage(socket *gws.Conn, message *gws.Message) {
}
