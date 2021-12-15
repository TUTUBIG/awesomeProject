package main

import "net/http/httputil"

func main() {
	mux := httputil.NewSingleHostReverseProxy()
	mux.BufferPool
}
