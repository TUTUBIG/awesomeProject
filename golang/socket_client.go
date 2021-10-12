package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

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
		Addr: [4]byte{127, 0, 0, 1},
	}

	if err := syscall.Connect(socket, &socketAddr); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("------connect-------socket: ", socket)

	f, e := os.Create("socket_test_file_1")
	if e != nil {
		fmt.Println(e)
		return
	}

	for {
		time.Sleep(time.Second)
		var buf [128]byte
		n, e := syscall.Read(socket, buf[:])
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
	return
}
