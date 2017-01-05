package scope

import "fmt"

var y int = 11

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	printLocalScope()

}

var z int = 12
