package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	bitLenSeq = 16
	bitLenSrv = 8

	shiftSec = bitLenSeq + bitLenSrv

	maskSeq   = uint32(-1 ^ (-1 << bitLenSeq))
	maskSrv   = uint64(-1 ^ (-1 << bitLenSrv))
	maskParse = uint64((1<<bitLenSeq - 1) << bitLenSrv)

	usUnit = int64(time.Millisecond)
)

type Generator struct {
	sec int64

	seq uint32
	srv uint8

	mu sync.Mutex
}

func (me *Generator) New() uint64 {
	me.mu.Lock()
	defer me.mu.Unlock()

	sec := time.Now().Unix()
	if me.seq <= maskSeq {
		me.seq = (me.seq + 1) & maskSeq
	} else {
		me.seq = 0
	}
	if me.sec == sec && me.seq == maskSeq {
		for sec <= me.sec {
			sec = time.Now().Unix()
		}
	}

	me.sec = sec

	fmt.Println(sec, me.seq, me.srv)

	return makeID(sec, me.seq, me.srv)
}

func makeID(sec int64, seq uint32, srv uint8) uint64 {
	return uint64(sec)<<shiftSec | uint64(seq)<<bitLenSrv | uint64(srv)
}

func main() {
	// g := new(Generator)
	//	fmt.Println(g.New())
	t := time.Now()

	fmt.Println(t.Unix(), t.UnixNano(), uint64(t.Unix())<<16|uint64(0))

	a := "/sdfsfds/fsdfsafds/bb/graphql/ui"
	i := strings.Index(a, "graphql")
	fmt.Println(a[i:])
}
