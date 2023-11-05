package main

import (
	"backend/api/server"
	"fmt"
	"os"
)

func main() {
	// Init context
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {

	// g := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	// Or instantiate a client with a custom bin and args options
	// 	llm, err := ollama.New(ollama.WithModel("llama2"))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	ctx := context.Background()

	// 	c.String(http.StatusOK, completion)
	// })

	server.InitializeServer()

	return nil
}
