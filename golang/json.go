package main

import (
	"encoding/json"
	"fmt"
)

type D struct {
	A string      `json:"a"`
	B interface{} `json:"b"`
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
}
