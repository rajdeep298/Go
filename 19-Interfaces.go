package main

import "math"

// Interface is a type that defines a set of methods.
// A value of interface type can hold any value that implements those methods.
type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.length * r.width
}

func (r rect) perimeter() float64 {
	return 2*r.length + 2*r.width
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (c circle) perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func measure(g geometry) {
	//println(g)
	println(g.area())
	println(g.perimeter()) // The type of g is determined by the underlying type of the value passed to it.

}
func main() {
	r := rect{length: 10, width: 5}
	c := circle{radius: 10}

	measure(r)
	measure(c)
}
