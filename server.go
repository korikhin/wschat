package main

import (
	"fmt"
	"os"
	"syscall"
)

func startServer() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating socket: %v\n", err)
		os.Exit(1)
	}

	addr := &syscall.SockaddrInet4{Port: 3000, Addr: [4]byte{127, 0, 0, 1}}
	err = syscall.Bind(fd, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error binding socket: %v\n", err)
		os.Exit(1)
	}

	err = syscall.Listen(fd, 10)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listening on socket: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Server is listening on port %d\n", addr.Port)
	for {
		nfd, _, err := syscall.Accept(fd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error accepting connection: %v\n", err)
			continue
		}
		go handleConnection(nfd)
	}
}

func handleConnection(fd int) {
	defer syscall.Close(fd)
	buf := make([]byte, 1024)
	for {
		n, err := syscall.Read(fd, buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from connection: %v\n", err)
			return
		}
		if n > 0 {
			message := buf[:n]
			_, err = syscall.Write(fd, message)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing to connection: %v\n", err)
				return
			}
			fmt.Println(string(message))
		}
	}
}
