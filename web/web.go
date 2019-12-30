package main

import (
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/a", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:63342")
		http.HandleFunc("/a/test", func(writer http.ResponseWriter, request *http.Request) {
			_, _ = writer.Write([]byte("hello"))
		})
	})
}

func main() {
	_ = http.ListenAndServe("127.0.0.1:9090", http.DefaultServeMux)
	time.Sleep(time.Hour)
}
