package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

const address = "https://apis.ankr.com/3a460696a4e043ce8022dccc5cb7ea03/7af9c7a27032e18bc95fda33ab009406/polygon/archive/main"

func main() {
	client, err := ethclient.Dial(address)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	// blockNumber := big.NewInt(18372872)
	ctx := context.Background()
	for {

		head, err := client.HeaderByNumber(ctx, nil)
		if err != nil {
			fmt.Println("got new block number: ", err)
			continue
		}

		number := new(big.Int).Sub(head.Number, big.NewInt(129))

		hash := fetchBlockFromApi(number.Uint64())

		if hash == "" {
			log.Fatal("error block ", number.Uint64())
		} else {
			log.Printf("good block,number %d,hash: %s", number.Uint64(), hash)
		}

		/*		officeHash := fetchFromOffice(number.Uint64())

				if officeHash == "" || hash == ""{
					log.Println("null hash")
					continue
				}
				if !strings.Contains(officeHash,hash) {
					log.Fatalf("block %d doesn't match. api hash %s, office hash: %s",number.Uint64(),hash,officeHash)
				} else {
					log.Printf("block %d hash %s match",number.Uint64(),hash)
				}*/

		time.Sleep(2 * time.Second)
	}
}

func fetchBlockFromApi(blockNumber uint64) string {
	res, err := http.Post(address, "application/json", strings.NewReader(fmt.Sprintf(`{"jsonrpc": "2.0", "id": "0", "method": "eth_getBlockByNumber", "params": ["0x%0x",false]}`, blockNumber)))
	if err != nil {
		log.Println(err)
		return ""
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	data := make(map[string]interface{})
	if err := json.Unmarshal(resBody, &data); err != nil {
		fmt.Println(err)
		return ""
	}
	result, ok := data["result"].(map[string]interface{})
	if !ok {
		return ""
	}
	hash, ok := result["hash"].(string)
	if !ok {
		return ""
	}
	return hash
}
