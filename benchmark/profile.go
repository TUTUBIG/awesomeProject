package main

import (
	"fmt"
	"math/rand"
	_ "net/http/pprof"
	"time"
)

func TestTime(i int) {
	fmt.Println("TestTime...")
	time.Sleep(time.Duration(i) * time.Second)
}

func TestMemery(i int) {
	fmt.Println("TestMemery...")
	a := make([][1024]byte, i, i)
	time.Sleep(time.Minute)
	fmt.Println(len(a))
}

func TestGoroutine(i int) {
	fmt.Println("TestGoroutine...")
	j := i
	for i > 0 {
		time.Sleep(10 * time.Microsecond)
		i--
		go func() {
			time.Sleep(time.Minute)
		}()
	}
	for j > 0 {
		j--
		go func() {
			time.Sleep(time.Second)
		}()
	}
}

func main() {

	//k := "121345"

	fmt.Printf("%03d", rand.Intn(1000))

	//fmt.Println(fmt.Sprintf("%c",int32(67)))

	/*// 26411491670361088
	id := int64(26412979238995968)

	n := make([]rune, 0)
	for i, r := range fmt.Sprintf("%d", id) {
		if i < 8 {
			continue
		}
		n = append(n, rune(r))
	}

	id, e := strconv.ParseInt(string(n), 10, 32)

	ss := int32(id)

	fmt.Println(id,e,ss)*/
	/*go TestTime(10)
	go TestMemery(50000)
	go TestGoroutine(5000)
	log.Println(http.ListenAndServe("localhost:6060",nil))*/
}
