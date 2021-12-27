package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set(`WWW-Authenticate`, `Basic realm="site"`)
		// writer.WriteHeader(http.StatusUnauthorized)
		//return
		writer.WriteHeader(http.StatusOK)
		log.Println(writer.Write([]byte("hello")))
		return
	})
	log.Fatal("start fail: ", http.ListenAndServe("0.0.0.0:8080", nil))
}
