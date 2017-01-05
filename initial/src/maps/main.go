package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	fmt.Println(m)
	_, exists := m["c"]
	fmt.Println(exists)


	delete(m, "c")
	fmt.Println(m["c"])
	fmt.Println(m)

	_, existsWithoutC := m["c"]
	fmt.Println(existsWithoutC)

	value, exists := m["a"]
	fmt.Println(value)

	newMap := map[string]int{"x":5,"y":10}
	fmt.Println(newMap)

	if value, exists := m["a"]; exists{
		fmt.Println(value)
	} else {
		fmt.Println("Ops")
	}
}
