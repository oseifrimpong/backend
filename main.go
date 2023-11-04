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

	"github.com/tmc/langchaingo/llms/ollama"
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

	r.GET("/ping", func(c *gin.Context) {
		// Or instantiate a client with a custom bin and args options
		llm, err := ollama.New(ollama.WithModel("llama2"))
		if err != nil {
			log.Fatal(err)
		}
		ctx := context.Background()
		completion, err := llm.Call(ctx, "Human: give me an example of a go code.\nAssistant:",
			llms.WithTemperature(0.8),
			llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				return nil
			}),
		)
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, completion)
	})

	// a post api for upload a pdf file and save to local
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		// save the file to local
		c.SaveUploadedFile(file, file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	r.Run(os.Getenv("APP_PORT"))

	return nil
}
