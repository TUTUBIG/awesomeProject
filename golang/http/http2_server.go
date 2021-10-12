package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})
	log.Fatal("start fail: ", http.ListenAndServeTLS("0.0.0.0:8080", "localhost8080.cert", "localhost8080.key", nil))
}
