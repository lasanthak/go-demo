package main

import (
	"fmt"
	"math"
)

// Define interfaces
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Drawable interface {
	Draw()
}

// Composite interface
type DrawableShape interface {
	Shape
	Drawable
}

// Implement interfaces implicitly
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Draw() {
	fmt.Printf("Drawing rectangle %v\n", r)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Draw() {
	fmt.Printf("Drawing circle with radius %.2f\n", c.Radius)
}

// Functions that accept interfaces
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func DrawAndMeasure(ds DrawableShape) {
	ds.Draw()
	PrintShapeInfo(ds)
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	// Interface values
	var shape Shape

	shape = rect
	fmt.Println("Rectangle:", shape)
	PrintShapeInfo(shape)

	shape = circle
	fmt.Println("Circle:", shape)
	PrintShapeInfo(shape)

	// Composite interface
	fmt.Println("\nUsing composite interface:")
	DrawAndMeasure(rect)
	DrawAndMeasure(circle)
}
