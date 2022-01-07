package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// raw()
	zip := compression("hello")
	unzip := decompression(zip)
	fmt.Println(string(unzip))
}

func raw() {
	b := new(bytes.Buffer)
	b.WriteString("hello")
	fmt.Println(b.String())
	r := make([]byte, 10)
	n, _ := b.Read(r)
	fmt.Println(string(r[:n]))
	r1 := make([]byte, 1)
	_, _ = b.Read(r1[:])
	fmt.Println(string(r1[:]))
}

func compression(origin string) []byte {
	b := new(bytes.Buffer)
	w := gzip.NewWriter(b)
	_, e := w.Write([]byte(origin))
	if e != nil {
		panic(e)
	}
	if e := w.Flush(); e != nil {
		panic(e)
	}

	if e := w.Close(); e != nil {
		panic(e)
	}

	return b.Bytes()
}

func decompression(zip []byte) []byte {
	r, e := gzip.NewReader(bytes.NewReader(zip))
	if e != nil {
		panic(e)
	}
	defer r.Close()

	/*unzip,e := ioutil.ReadAll(r)
	if e != nil {
		panic(e)
	}*/

	var unzip [128]byte
	n, e := r.Read(unzip[:])
	if e != nil {
		panic(e)
	}

	return unzip[:n]
}

func standardHttp() {
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
