package main

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	res, err := http.Get("http://baidu.com")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	buffer := bytes.NewBuffer([]byte{})
	if _, err := io.Copy(buffer, res.Body); err != nil {
		log.Fatal(err)
	}

	log.Println("resp", buffer.String())

	fastHttp()
}

func fastHttp() {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	req.SetRequestURI("http://baidu.com")
	req.Header.SetMethod("GET")

	c := fasthttp.Client{
		NoDefaultUserAgentHeader:      true,
		MaxConnsPerHost:               1000,
		ReadBufferSize:                6144, // 6Kib
		WriteBufferSize:               6144, // 6Kib
		ReadTimeout:                   30 * time.Second,
		WriteTimeout:                  30 * time.Second,
		MaxIdleConnDuration:           3 * time.Minute,
		DisableHeaderNamesNormalizing: true,
	}
	if e := c.Do(req, resp); e != nil {
		log.Fatal(e)
	}

	buffer := bytes.NewBuffer([]byte{})
	if _, err := resp.WriteTo(buffer); err != nil {
		log.Fatal(err)
	}
	log.Println("fast http resp", buffer.String())
}
