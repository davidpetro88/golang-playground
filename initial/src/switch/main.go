package main

import "fmt"

func main() {
	getName("Mitchel")
	getName("David")
	getName("Phil")
}

func getName( name string)  {
	switch name {
	case "David":
		fmt.Println("Hey ", name)
	case "Phil":
		fmt.Println("Hey Phil")
	default:
		fmt.Println("I did not find your name! ", name)
	}
}
