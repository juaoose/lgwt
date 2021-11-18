package shapes

import "math"

// Interface resolution is implicit:
// if the interface is satisfied by a type, then the type is used
// Both rectangle and circle have an implementation of the Area method
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Pr = 2 * (w + h)
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Pc = 2 * π * r
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Pt = 2 * (w + h)
func (t Triangle) Perimeter() float64 {
	return t.Base + t.Base + t.Height
}

// Ar = w * h
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Ac = π * r^2
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// At = (b * h) / 2
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
