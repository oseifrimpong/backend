package main

import (
	"backend/api/server"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {
	server.InitializeServer()
	return nil
}
