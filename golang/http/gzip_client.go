package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	data := `{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":0}`
	b := new(bytes.Buffer)
	w := gzip.NewWriter(b)
	w.Write([]byte(data))
	w.Flush()
	w.Close()

	req, _ := http.NewRequest(http.MethodGet, "https://apis-stage.ankr.com/b2870b2a3ec448008017a4e3b0d01b3c/057b1fcba4f9894b308e1e11c11864a7/binance/full/main", b)
	req.Header.Set("Content-Encoding", "gzip")
	// req.Header.Set("Content-Type","application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	if res.Header.Get("Content-Encoding") == "gzip" {
		r, _ := gzip.NewReader(res.Body)
		data, _ := ioutil.ReadAll(r)
		log.Println("--data raw--zip", string(data))
		r.Close()
		res.Body.Close()
		return
	}

	datares, _ := ioutil.ReadAll(res.Body)
	log.Println("--data raw--", string(datares))
	res.Body.Close()
	return
}
