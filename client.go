package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
)

const exitCommand = "exit"

func startClient() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating socket: %v\n", err)
		os.Exit(1)
	}

	defer func() {
		syscall.Close(fd)
		fmt.Println("Connection closed")
	}()

	addr := &syscall.SockaddrInet4{Port: 3000, Addr: [4]byte{127, 0, 0, 1}}
	err = syscall.Connect(fd, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting: %v\n", err)
		os.Exit(1)
	}

	console := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a messages, type %q to quit:\n", exitCommand)

	for console.Scan() {
		input := console.Text()
		if input == "" {
			continue
		}
		if input == "exit" {
			break
		}

		_, err := syscall.Write(fd, []byte(input))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading the input: %v\n", err)
			syscall.Close(fd)
			break
		}

		buf := make([]byte, 1024)
		n, err := syscall.Read(fd, buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading the input: %v\n", err)
			syscall.Close(fd)
			break
		}
		fmt.Println("Echo:", string(buf[:n]))
	}
}
