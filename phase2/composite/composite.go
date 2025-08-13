package main

import (
	"fmt"
)

// Composition vs inheritance
type Engine struct {
	Horsepower int
	Type       string
}

func (e Engine) Start() {
	fmt.Printf("Starting %s engine (%d HP)\n", e.Type, e.Horsepower)
}

func (e Engine) Stop() {
	fmt.Printf("Stopping %s engine\n", e.Type)
}

type GPS struct {
	CurrentLocation string
}

func (g GPS) Navigate(destination string) {
	fmt.Printf("Navigating from %s to %s\n", g.CurrentLocation, destination)
}

type Car struct {
	Brand  string
	Model  string
	Engine Engine // Composition - Car HAS-A Engine
	GPS    GPS    // Composition - Car HAS-A GPS
}

func (c Car) Drive(destination string) {
	c.Engine.Start()
	c.GPS.Navigate(destination)
	fmt.Printf("Driving %s %s to %s\n", c.Brand, c.Model, destination)
}

func (c Car) Park() {
	c.Engine.Stop()
	fmt.Printf("Parked %s %s\n", c.Brand, c.Model)
}

// Decorator pattern using composition
type Coffee interface {
	Cost() float64
	Description() string
}

type SimpleCoffee struct{}

func (sc SimpleCoffee) Cost() float64 {
	return 2.00
}

func (sc SimpleCoffee) Description() string {
	return "Simple coffee"
}

type MilkDecorator struct {
	Coffee Coffee // Composition
}

func (md MilkDecorator) Cost() float64 {
	return md.Coffee.Cost() + 0.50
}

func (md MilkDecorator) Description() string {
	return md.Coffee.Description() + ", milk"
}

type SugarDecorator struct {
	Coffee Coffee // Composition
}

func (sd SugarDecorator) Cost() float64 {
	return sd.Coffee.Cost() + 0.25
}

func (sd SugarDecorator) Description() string {
	return sd.Coffee.Description() + ", sugar"
}

// Adapter pattern
type LegacyPrinter struct{}

func (lp LegacyPrinter) PrintOldFormat(text string) {
	fmt.Printf("[LEGACY] %s\n", text)
}

type ModernPrinter interface {
	Print(text string)
}

type PrinterAdapter struct {
	LegacyPrinter LegacyPrinter // Composition
}

func (pa PrinterAdapter) Print(text string) {
	pa.LegacyPrinter.PrintOldFormat(text)
}

// Observer pattern
type EventPublisher struct {
	subscribers []EventSubscriber
}

type EventSubscriber interface {
	Notify(event string)
}

func (ep *EventPublisher) Subscribe(subscriber EventSubscriber) {
	ep.subscribers = append(ep.subscribers, subscriber)
}

func (ep *EventPublisher) Publish(event string) {
	for _, subscriber := range ep.subscribers {
		subscriber.Notify(event)
	}
}

type EmailNotifier struct {
	Email string
}

func (en EmailNotifier) Notify(event string) {
	fmt.Printf("Email to %s: %s\n", en.Email, event)
}

type SMSNotifier struct {
	PhoneNumber string
}

func (sn SMSNotifier) Notify(event string) {
	fmt.Printf("SMS to %s: %s\n", sn.PhoneNumber, event)
}

func main() {
	// Car composition example
	car := Car{
		Brand: "Toyota",
		Model: "Camry",
		Engine: Engine{
			Horsepower: 200,
			Type:       "V6",
		},
		GPS: GPS{
			CurrentLocation: "Home",
		},
	}

	car.Drive("Work")
	car.Park()

	fmt.Println("---")

	// Decorator pattern
	coffee := SimpleCoffee{}
	fmt.Printf("%s: $%.2f\n", coffee.Description(), coffee.Cost())

	coffeeWithMilk := MilkDecorator{Coffee: coffee}
	fmt.Printf("%s: $%.2f\n", coffeeWithMilk.Description(), coffeeWithMilk.Cost())

	coffeeWithMilkAndSugar := SugarDecorator{Coffee: coffeeWithMilk}
	fmt.Printf("%s: $%.2f\n", coffeeWithMilkAndSugar.Description(), coffeeWithMilkAndSugar.Cost())

	fmt.Println("---")

	// Adapter pattern
	var printer ModernPrinter = PrinterAdapter{
		LegacyPrinter: LegacyPrinter{},
	}
	printer.Print("Hello, World!")

	fmt.Println("---")

	// Observer pattern
	publisher := &EventPublisher{}

	emailNotifier := EmailNotifier{Email: "user@example.com"}
	smsNotifier := SMSNotifier{PhoneNumber: "+1234567890"}

	publisher.Subscribe(emailNotifier)
	publisher.Subscribe(smsNotifier)

	publisher.Publish("New order received!")
	publisher.Publish("Payment processed!")
}
