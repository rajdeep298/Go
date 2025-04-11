package main

type rectangle struct {
	length int
	width  int
}

func (r *rectangle) area() int {
	return r.length * r.width
}

func (r rectangle) perimeter() int {
	return 2*r.length + 2*r.width
}

func main() {
	r := rectangle{length: 10, width: 5}
	println("Area: ", r.area())
	println("Perimeter: ", r.perimeter())

	rp := &r
	println("Area: ", rp.area())
	println("Perimeter: ", rp.perimeter())
}
