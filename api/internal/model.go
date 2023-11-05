package internal

import (
	"log"

	"github.com/tmc/langchaingo/llms/ollama"
)

func NewLLM() (*ollama.LLM, error) {

	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		log.Fatal(err)
	}
	return llm, nil
}
