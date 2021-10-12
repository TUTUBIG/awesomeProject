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
		Addr: [4]byte{127, 0, 0, 1},
	}

	if err := syscall.Bind(socket, &socketAddr); err != nil {
		log.Fatal(err)
	}

	if err := syscall.Listen(socket, 20); err != nil {
		log.Fatal(err)
	}

	fmt.Println("socket: ", socket)

	for {
		nfd, _, err := syscall.Accept(socket)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("-----accept---new socket: ", nfd)

		go func() {
			f, e := ioutil.ReadFile(file)
			if e != nil {
				log.Fatal(e)
			}

			_, e = syscall.Write(nfd, f)
			if e != nil {
				log.Fatal(err)
			}
		}()

	}
}
