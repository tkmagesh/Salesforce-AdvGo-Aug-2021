package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (s *MockService) Send(msg string) bool {
	fmt.Println("Mock Send() method invoked")
	args := s.Called(msg)
	return args.Bool(0)
}

func TestMessageProcessor(t *testing.T) {
	mockService := &MockService{}
	mockService.On("Send", "Hello").Return(true)

	processor := MessageProcessor{mockService}
	processor.Process("Hello")

	mockService.AssertExpectations(t)
}
