package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { // range iterates over elements in a variety of data structures.
		sum += num
	}
	fmt.Println("Sum:", sum)

	for i, num := range nums { // range on arrays and slices provides both the index and value for each entry.
		if num == 3 {
			fmt.Println("Index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"} // range on map iterates over key/value pairs.
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs { // range can also iterate over just the keys of a map.
		fmt.Println("key:", k)
	}

	for i, c := range "go" { // range on strings iterates over Unicode code points.
		fmt.Println(i, c)
	}
}
