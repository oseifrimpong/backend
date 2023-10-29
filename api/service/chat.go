package service

type ChatService interface {
	QueryChat() error
}

type chatSvc struct{}

func NewChatService() ChatService {
	return &chatSvc{}
}

func (c *chatSvc) QueryChat() error {
	return nil
}
