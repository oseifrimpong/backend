package service

import (
	"backend/api/internal"
	"context"
	"log"

	"github.com/tmc/langchaingo/llms"
)

type ChatService interface {
	QueryChat(question string) (string, error)
}

type chatSvc struct {
	ctx context.Context
}

func NewChatService() ChatService {
	ctx := context.Background()
	return &chatSvc{ctx: ctx}
}

func (c *chatSvc) QueryChat(question string) (string, error) {

	llm, err := internal.NewLLM()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	completion, err := llm.Call(c.ctx, question,
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			return nil
		}),
	)

	if err != nil {
		log.Fatal(err)
	}
	return completion, nil
}
