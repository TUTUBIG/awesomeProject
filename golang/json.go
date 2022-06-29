package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

type D struct {
	A string        `json:"a"`
	B interface{}   `json:"b"`
	C []interface{} `json:"c"`
	D []string      `json:"d"`
}

func main() {
	d := `{"A":"a"}`
	dd := new(D)
	_ = json.Unmarshal([]byte(d), dd)
	switch dd.B.(type) {
	case string:
		fmt.Println("string")
	case []interface{}:
		fmt.Println("[]interface")
	case interface{}:
		fmt.Println("interface")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("default")
	}
	fmt.Println(dd)

	d1 := D{
		A: "",
		B: nil,
		C: nil,
		D: nil,
	}

	data1, _ := json.Marshal(d1)
	fmt.Println(string(data1))

	d1 = D{
		A: "",
		B: nil,
		C: make([]interface{}, 0),
		D: make([]string, 0),
	}

	data1, _ = json.Marshal(d1)
	fmt.Println(string(data1))
	b := `{
 "parentHash": "0x0f5fe54d0242423bfe83d448a29031235b1dc89a68ee3725bcf4ae420e9b8fad",
 "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
 "miner": "0xac0e15a038eedfc68ba3c35c73fed5be4a07afb5",
 "stateRoot": "0xc407a5c4e52501b377bce19ea74a6cd84b66554be139f73284bcb3233d918f35",
 "transactionsRoot": "0xbe1f7350d76ba085442cd65a1b90976ebe1cc3660e29cffc94f52e9ec4c48aa0",
 "receiptsRoot": "0x0af4ff5171bdb4616179c3b41f34a23d4ae017e4b5817665e67a810b4bf782e3",
 "logsBloom": "0xfdf9a77f648cfb5d1f6d66bdcd9e7afffedafbddb7f5847f7ddef3bfef7ed6fc6b937f5ef87efddfbb6d7eff6f13eeacfbbbf4d9ff3bed6978fe1e751eaff7fea4efacaeda7b4b66e5e3fbede7ff77f83fdfff3bdfc7ffbfafb5c6b7f23577375fdd7ef5bbfee7f6df7d3edf79cbeb7d8953f6df9eb9aefebea00e3bfbffb7f4ff63e7ade67fff9a8ef55f0fde3a5fee79ef9ef7bfe3fbffe9dfc1d9ad7f7d3bda7e7323dd7bf8ef7419ed6ffeeff283f5df2d2d7adbefdf7a8dff1ff69fecefbfb95f32bfeb27ffb5de35fd6df3f7eb7d7b676eff1a7e1ef9dcb1efcabaf5ebfe7ffb6f555df65e7f29374f1f6c7cdef7af9f7fe0dedffdee92fffaffd3aff5",
 "difficulty": "0x2",
 "number": "0xe937f4",
 "gasLimit": "0x4bf72eb",
 "gasUsed": "0x4b800d8",
 "timestamp": "0x620beb85",
 "extraData": "0xd983010107846765746889676f312e31362e3130856c696e75780000c3167bdf886a144c012a75ba2b00bd9a1b118898e4bd91779be0e5b7dcea65fab331fc6b4d1a2ec8e0d21ccee90d32db8be7fb5167c01292c7179931e94fa7dd5f11d49901",
 "mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
 "nonce": "0x0000000000000000",
 "hash": "0xe6cdad5b092ac763c162b3e0caa1f6a6b396a70620e006d908e554e84b3c1d8d"
}`
	// b := `{"difficulty":"0x2","extraData":"0xd883010107846765746888676f312e31362e37856c696e7578000000c3167bdf09a41f0a18ebda0b6c95a7e1d7b4c6b94e3f2ade1ef88660de0a83cc6ec75d8d7c5a6461aaf47bbe866b73b35281a9feb52d44fa923cee1cd80bbf7451b0150c00","gasLimit":"0x4bee2ba","gasUsed":"0x26cc400","hash":"0xab94c76e01dd967737f256560b7d518e1adea809f8dc3e88b7e8055a8b3f05a1","logsBloom":"0xbde67b6da467d6fb97f8d4a7eba7168e2c68afcf3467e2fddd14b5687e2665704cfa7fd768de5d3ce9af3afe2cbf17736edfbf11ef3e5f7f7dbc8fdb77fd25e8f466a3b737e75b726d7b5bf9ff3373fefffc32dffacdc87ff4afdffd91bd86b5c8debffdf66f70936b7567fd51ed5be5bbd6bfdfb0be66df2670c35d3b4dd6af1ed8abafcf9dbd453da3bd137f9cc3e04876b625ffffd61fff4fbb655dfb1fb682f7deed4be2cdcf4ed6756dab8e8cbc778c4ae5f1ebfa7394bb2f75dafa6cedefc7b71b20fb68a355f7af9f3ef2fef3b76e37f79fc73f5f8a7ff8fef0f7f7e0fb371ed120515eebdfcc5fb5c95def00e7ffec4f41f1f66fb553fbfefd9bf8b7","miner":"0xEF0274E31810C9Df02F98FAFDe0f841F4E66a1Cd","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","number":"0xe9713d","parentHash":"0xd1ce26a82e630e03065792bad561758916a08980b2583a6a610d6a26be6d1de7","receiptsRoot":"0xdad5abccce5cc540f83ff033089f7d26ca3b8289486832cc3f6b542f10d06857","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","stateRoot":"0x329b72b8dedb1569b076b83b881f9901c347eefcb358e9d8c07fc73b02766681","timestamp":"0x620c9834","transactionsRoot":"0x758738fd72d9bb741c6cf117188fd4efc22bc627ad3fb9807a1f9fdfcdc472a0"}`
	// b := `{"parentHash":"0xbff83c998fc99236ad5121a29f2f912956f14f8fea9ce0e5f997990280ce293f","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","miner":"0x70f657164e5b75689b64b7fd1fa275f334f28e18","stateRoot":"0x3a050b22ce3f974be70741b4eebcca4e4be415185e2dc2ccf3067437007c1091","transactionsRoot":"0x76ebf0177a8f4a9bef11ec0972e08c771f8b16f9b503f81069d535589c6f415d","receiptsRoot":"0xf090a5006da41c7a358aa7be7edf6f2b6e878b27133586759daea55f48bc564d","logsBloom":"0xdfffff76efef73f7f6ede6ffff563fe6f9d7bcffff7dfeffffff5e73fffb5f35ee67bffff7ef9fd7ffef5efcafffabbabfffdeb5fefeef59ffae4f7c7eef77ebfbffef2a77ff3f6d4fffffffaf7ffff3bffb4fedbbf7fdeebff4efd7fcb8fddf6ffdfff7fafbfffd33d6f7eb7fffff7dff0efffbff5defef7fffdffdf6b7fffeffffe71ff73ff1bfd6fdfff3fff72ffd99f77fbffff7dcfffffd6e7b6efd9df3ff7ff7bfddfffeec7ddffdd3f37efe7be7e5fffdffdefbaddef77bffbefefffbfbf77fee7bdef7ffdafffff7fdaf7ffbfffcfef6bffde5fdd3ffffdfbff2ffa7f2faf9796ed95dfd5faafffff5ff5d5df5fffdfb777fbf5fbfbbfdfdedff17db","difficulty":"0x2","number":"0xe97198","gasLimit":"0x4c4b400","gasUsed":"0x3c3542d","timestamp":"0x620c9945","extraData":"0xd883010107846765746888676f312e31332e34856c696e7578000000c3167bdfc75597dbf135c32bb91c3d41e9a10e274510761c181b682afc9a7126991e55c923c700e1a6de7e98767e87de7fec65687430a69a33c813bc695a1b4b192e44bf01","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","hash":"0xed3c70efb75f5a342f5c516f23a28804fe2747f988aeb2d17d6639f223f10b14"}`
	h := new(types.Header)
	bb, ee := h.MarshalJSON()
	fmt.Println(string(bb), ee)
	e := json.Unmarshal([]byte(b), h)
	fmt.Println(e)

	r := `{"id":1}`
	r1 := `{"id":1.1}`
	r2 := `{"id":"1.2a"}`
	type R struct {
		Id json.RawMessage `json:"id"`
	}

	var rr, rr1, rr2 R
	if err := json.Unmarshal([]byte(r), &rr); err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(r1), &rr1); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal([]byte(r2), &rr2); err != nil {
		log.Fatal(err)
	}

	if aa, err := rr.Id.MarshalJSON(); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(aa))
		log.Printf("aa%+v", aa)
	}

	if aa, err := rr1.Id.MarshalJSON(); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(aa))
		log.Printf("aa%+v", aa)
	}

	if aa, err := rr2.Id.MarshalJSON(); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(aa))
		log.Printf("aa%v", aa)
	}

	if dd, err := json.Marshal(rr); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(dd))
	}

	if dd, err := json.Marshal(rr1); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(dd))
	}

	if dd, err := json.Marshal(rr2); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(dd))
	}

}
