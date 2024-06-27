package perimeter

import "math"

type Shape interface {
	GetArea() float64
	GetPerimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) GetPerimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) GetArea() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) GetArea() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) GetPerimeter() float64 {
	return 2 * math.Pi * c.Radius
}
