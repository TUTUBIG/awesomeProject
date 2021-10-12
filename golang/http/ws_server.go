package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upGrader = &websocket.Upgrader{}

func main() {
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		websocketConnection, err := upGrader.Upgrade(writer, request, nil)
		if err != nil {
			panic(err)
		}

		websocketConnection.ReadMessage()

		for {
			time.Sleep(time.Second)
			if err := websocketConnection.WriteJSON(1); err != nil {
				panic(err)
			}
		}
	})
	log.Fatal("start fail: ", http.ListenAndServe("0.0.0.0", nil))
}
