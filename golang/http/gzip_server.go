package main

import (
	"compress/gzip"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		ae := request.Header.Get("Accept-Encoding")
		vv := strings.Split(ae, ",")
		for _, v := range vv {
			if v == "gzip" {
				writer.Header().Set("Content-Encoding", "gzip")
				w := gzip.NewWriter(writer)
				log.Println("gzip")
				log.Println(w.Write([]byte("test hello11")))
				w.Close()
				return
			}
		}

		log.Println("raw")
		log.Println(writer.Write([]byte("test hello")))
	})
	log.Println(http.ListenAndServe(":8088", http.DefaultServeMux))
}
