package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/local"
)

func main() {
	// Init context
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run(ctx context.Context) error {

	// g := gin.Default()
	r := gin.Default()

	switch gin.Mode() {
	case gin.ReleaseMode:
		// logger = config.Logger()
	default:
		// logger = config.Logger()

		err := godotenv.Load()
		if err != nil {
			fmt.Println("error loading .env file")
		}
	}

	// You may instantiate a client with a default bin and args from environment variable
	// llm, err := local.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Or instantiate a client with a custom bin and args options
	clientOptions := []local.Option{
		local.WithBin(os.Getenv("LOCAL_LLM_BIN")),
		// local.WithArgs("--arg1=value1 --arg2=value2"),
		local.WithGlobalAsArgs(), // build key-value arguments from global llms.Options, then append to args
	}
	llm, err := local.New(clientOptions...)

	// By default, library will use default bin and args
	// completion, err := llm.Call(ctx, "How many sides does a square have?")
	// Or append to default args options from global llms.Options
	generateOptions := []llms.CallOption{
		// llms.WithTopK(10),
		// llms.WithTopP(0.95),
		// llms.WithSeed(13),
		llms.WithModel("llama"),
	}
	// In that case command will look like: /path/to/bin --arg1=value1 --arg2=value2 --top_k=10 --top_p=0.95 --seed=13 "How many sides does a square have?"
	completion, err := llm.Call(ctx, "How many sides does a square have?", generateOptions...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, completion)
	})
	r.Run(os.Getenv("APP_PORT"))

	return nil
}
