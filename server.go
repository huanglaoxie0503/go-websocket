package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go-websocket/impl"
	"net/http"
	"time"
)

var (
	upgrade = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// 回调函数
func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		conn   *impl.Connection
		err    error
		data   []byte
	)
	if wsConn, err = upgrade.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	// 启动一个协程，每隔1秒发送一次消息
	go func() {
		for {
			if err := conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}
ERR:
	// TODO:关闭websocket连接操作
}

func main() {
	// http://localhost:8080/ws
	http.HandleFunc("/ws", wsHandler)

	// 监听端口
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
