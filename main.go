package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [server|client]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "server":
		startServer()
	case "client":
		startClient()
	default:
		fmt.Println("Invalid command. Use 'server' or 'client'")
		os.Exit(1)
	}
}
