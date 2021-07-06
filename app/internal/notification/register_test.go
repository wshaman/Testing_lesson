package notification_test

import (
	"fmt"
	"gtest_example/app/internal/notification"
	"testing"

	"github.com/stretchr/testify/mock"
)

type emlServiceMock struct {
	mock.Mock
}

func (m emlServiceMock) NotifyOnRegister(value string) (string, error) {
	fmt.Println("Mocked NotifyOnRegister")
	fmt.Printf("Value passed in: %s\n", value)
	args := m.Called(value)
	return args.String(0), args.Error(1)
}

func TestRegisterClient(t *testing.T) {
	emlService := new(emlServiceMock)

	emlService.On("NotifyOnRegister", "alice@example.com").
		Return("123", nil)

	// next we want to define the service we wish to test
	myService := notification.NewMessenger(emlService)
	c := notification.Client{
		Address: "alice@example.com",
	}
	myService.Register(c)

	emlService.AssertExpectations(t)
}
