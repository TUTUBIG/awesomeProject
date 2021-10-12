package main

import (
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

var upgrader = websocket.FastHTTPUpgrader{}

func wsHandler(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("test error")
	ctx.Error("test error", websocket.CloseNormalClosure)
	return

	err := upgrader.Upgrade(ctx, func(ws *websocket.Conn) {
		for {
			time.Sleep(time.Second)
			log.Println("", ws.WriteJSON("test"))
		}
	})

	if err != nil {
		if _, ok := err.(websocket.HandshakeError); ok {
			log.Fatal(err)
		}
		return
	}
}

func main() {
	server := &fasthttp.Server{
		Name:                  "test",
		Handler:               wsHandler,
		NoDefaultServerHeader: true,
		NoDefaultDate:         true,
		NoDefaultContentType:  true,
		DisableKeepalive:      false,
		TCPKeepalive:          true,
		ReadBufferSize:        4096,
		WriteBufferSize:       4096,
		ReadTimeout:           300 * time.Second,
		WriteTimeout:          300 * time.Second,
		IdleTimeout:           300 * time.Second,
	}
	if err := server.ListenAndServe(":8080"); err != nil {
		log.Fatal(err)
	}
}
