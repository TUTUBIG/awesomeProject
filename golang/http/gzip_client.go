package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// data := `{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":0}`
	f, e := ioutil.ReadFile(`eth_call_bsc.json`)
	if e != nil {
		panic(e)
	}

	// log.Println(len(f))

	b := new(bytes.Buffer)
	w := gzip.NewWriter(b)

	w.Write(f)
	w.Flush()
	w.Close()

	// b.Write(f)

	log.Println(b.Len())

	req, _ := http.NewRequest(http.MethodPost, "https://apis-stage.ankr.com/d468fe2166e94ac6b8062a89909a9870/7af9c7a27032e18bc95fda33ab009406/binance/full/main", b)
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
