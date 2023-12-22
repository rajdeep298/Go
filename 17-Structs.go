package main

import "fmt"

type person struct {
	name string
	age  int
}

func structure(x string) *person {
	person1 := person{name: x}
	person1.age = 42
	return &person1
}
func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	s2 := structure("John")
	fmt.Println(s2)

	sp := &s // Structs are mutable.
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name string
		age  int
	}{
		name: "Dog",
		age:  10,
	} //{
	//	"Dog",
	//	10,
	//}
	fmt.Println(dog)
}
