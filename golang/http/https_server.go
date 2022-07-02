package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		log.Println(writer.Write([]byte("hello")))
	})
	log.Fatal("start fail: ", http.ListenAndServeTLS("0.0.0.0:443", "keystore.pem", "keystore.key", nil))
}
