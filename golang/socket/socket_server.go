package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"syscall"
)

const file = "/Users/liukui/go/src/awesomeProject/socket_test_file"

func main() {
	socket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := syscall.Close(socket); err != nil {
			log.Fatal(err)
		}
	}()

	socketAddr := syscall.SockaddrInet4{
		Port: 31028,
		Addr: [4]byte{192, 168, 3, 35},
	}

	if err := syscall.Bind(socket, &socketAddr); err != nil {
		log.Fatal(err)
	}

	if err := syscall.Listen(socket, 3); err != nil {
		log.Fatal(err)
	}

	fmt.Println("listen socket: ", socket)

	for {
		nfd, sa, err := syscall.Accept(socket)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("-----accept---new socket: ", nfd, sa)

		go func(fd int) {
			pong(fd)
		}(nfd)
	}
}

func pong(fd int) {
	body := make([]byte, 128)
	for {
		n, e := syscall.Read(fd, body)
		if e != nil {
			log.Println(e)
		} else {
			log.Println("Pong: ", string(body[:n]))
		}
	}
}

func sendFile(fd int) {
	f, e := ioutil.ReadFile(file)
	if e != nil {
		log.Fatal(e)
	}

	_, e = syscall.Write(fd, f)
	if e != nil {
		log.Fatal(e)
	}
}
