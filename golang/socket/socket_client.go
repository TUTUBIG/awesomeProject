package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

const file1 = "/Users/liukui/go/src/awesomeProject/socket_test_file1"

func main() {
	socket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := syscall.Close(socket); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("client end")
	}()

	socketAddr := syscall.SockaddrInet4{
		Port: 31028,
		Addr: [4]byte{192, 168, 3, 35},
	}

	if err := syscall.Connect(socket, &socketAddr); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("------connect-------socket: ", socket)

	go func(fd int) {
		ping(fd)
	}(socket)

	time.Sleep(time.Minute)
	return
}

func ping(fd int) {
	for true {
		time.Sleep(time.Second)
		_, e := syscall.Write(fd, []byte("Ping"))
		if e != nil {
			log.Println(e)
		}
	}

}

func readFile(fd int) {
	f, e := os.Create(file1)
	if e != nil {
		fmt.Println(e)
		return
	}

	for {
		time.Sleep(time.Second)
		var buf [128]byte
		n, e := syscall.Read(fd, buf[:])
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(n)

		n, e = f.Write(buf[:n])
		if e != nil {
			fmt.Println(e)
			return
		}

	}

	fmt.Println(f.Sync(), f.Close())
}
