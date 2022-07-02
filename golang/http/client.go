package main

import (
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://localhost.com/test")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.StatusCode)
}
