package services

import "fmt"

type MessageService interface {
	Send(msg string) bool
}

type SMSService struct {
}

func (s *SMSService) Send(msg string) bool {
	return true
}

type MessageProcessor struct {
	MessageService MessageService
}

func (m *MessageProcessor) Process(msg string) bool {
	fmt.Println("processing the message using the given service")
	return m.MessageService.Send(msg)
}
