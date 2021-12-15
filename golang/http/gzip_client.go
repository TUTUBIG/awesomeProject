package main

import (
	"bytes"
	"compress/gzip"
	"log"
	"net/http"
)

func main() {
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:8088/test", bytes.NewReader([]byte{}))
	res, _ := http.DefaultClient.Do(req)

	log.Println("====", res.Header.Get("Content-Encoding"))

	if res.Header.Get("Content-Encoding") == "gzip" {
		r, _ := gzip.NewReader(res.Body)
		var data []byte
		log.Println(r.Read(data))
		log.Println("--data raw--zip", data)
	}
}
