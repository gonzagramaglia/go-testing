package shapes

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

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) GetArea() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) GetPerimeter() float64 {
	return 0
}

/*

"This was an important chapter
because we are now starting to define our own types.
In statically typed languages like Go,
being able to design your own types is essential
for building software that is easy to understand,
to piece together and to test.

Interfaces are a great tool for hiding complexity
away from other parts of the system.
In our case our test helper code did not need to
know the exact shape it was asserting on, only how
to "ask" for its area.

As you become more familiar with Go you will start to
see the real strength of interfaces and the standard
library. You'll learn about interfaces defined in
the standard library that are used everywhere and by
implementing them against your own types, you can very
quickly re-use a lot of great functionality.

*/
