package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	s := `0xd983010106846765746889676f312e31362e3130856c696e75780000c3167bdfdbf43e807a120c6602dac1423bb89b4a255fff72aee0ad175c5ce2db49bcb0ef1470f12018e76fa1e7bb84145f9922293bb06239ec644ddf187002c5ca0cc31b01`

	h := hexutil.Bytes(s)

	fmt.Println(string(h), h.String())
}
