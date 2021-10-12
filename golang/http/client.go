package main

import (
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.StatusCode)
}
