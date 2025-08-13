package main

import (
	"fmt"
	"io"
	"strings"
)

// Common Go interfaces
type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

type ReadWriter interface {
	Reader
	Writer
}

// Interface segregation
type Database interface {
	Connect() error
	Close() error
}

type QueryExecutor interface {
	Query(sql string) ([]Row, error)
}

type CommandExecutor interface {
	Execute(sql string) error
}

type FullDatabase interface {
	Database
	QueryExecutor
	CommandExecutor
}

// Strategy pattern with interfaces
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type CreditCardProcessor struct {
	CardNumber string
}

func (ccp CreditCardProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("Processing $%.2f payment via credit card ending in %s\n",
		amount, ccp.CardNumber[len(ccp.CardNumber)-4:])
	return nil
}

type PayPalProcessor struct {
	Email string
}

func (pp PayPalProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("Processing $%.2f payment via PayPal account %s\n", amount, pp.Email)
	return nil
}

type PaymentService struct {
	processor PaymentProcessor
}

func (ps *PaymentService) SetProcessor(processor PaymentProcessor) {
	ps.processor = processor
}

func (ps *PaymentService) ProcessOrder(amount float64) error {
	return ps.processor.ProcessPayment(amount)
}

// Interface for testing/mocking
type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Printf("[LOG] %s\n", message)
}

type MockLogger struct {
	Messages []string
}

func (ml *MockLogger) Log(message string) {
	ml.Messages = append(ml.Messages, message)
}

type Service struct {
	logger Logger
}

func (s *Service) DoWork() {
	s.logger.Log("Starting work")
	// Do some work...
	s.logger.Log("Work completed")
}

func main() {
	// io.Reader example
	text := "Hello, Go interfaces!"
	reader := strings.NewReader(text)

	buffer := make([]byte, 10)
	n, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading: %v\n", err)
	}
	fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))

	// Strategy pattern
	paymentService := &PaymentService{}

	// Use credit card processor
	creditCard := CreditCardProcessor{CardNumber: "1234567890123456"}
	paymentService.SetProcessor(creditCard)
	paymentService.ProcessOrder(99.99)

	// Switch to PayPal processor
	paypal := PayPalProcessor{Email: "user@example.com"}
	paymentService.SetProcessor(paypal)
	paymentService.ProcessOrder(149.99)

	// Logger example
	service := &Service{logger: ConsoleLogger{}}
	service.DoWork()

	// Mock logger for testing
	mockLogger := &MockLogger{}
	service.logger = mockLogger
	service.DoWork()
	fmt.Printf("Mock captured: %v\n", mockLogger.Messages)
}

type Row struct {
	Data map[string]interface{}
}
