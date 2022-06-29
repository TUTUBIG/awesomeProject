package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		vv := strings.Split(request.Header.Get("Content-Encoding"), ",")
		for _, v := range vv {
			if v == "gzip" {
				r, e := gzip.NewReader(request.Body)
				if e != nil {
					panic(e)
				}

				d, e := ioutil.ReadAll(r)
				if e != nil {
					panic(e)
				}
				fmt.Println("gzip data", string(d))
			}
		}

		vv = strings.Split(request.Header.Get("Accept-Encoding"), ",")
		for _, v := range vv {
			if v == "gzip" {
				writer.Header().Set("Content-Encoding", "gzip")
				w := gzip.NewWriter(writer)
				log.Println("gzip")
				log.Println(w.Write([]byte("hello server gzip")))
				_ = w.Flush()
				_ = w.Close()
				return
			}
		}

		log.Println("raw")
		log.Println(writer.Write([]byte("hello server")))
	})
	log.Println(http.ListenAndServe(":8088", http.DefaultServeMux))
}
