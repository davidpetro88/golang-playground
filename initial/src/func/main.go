package main

import "fmt"

func main() {
	x := funcName(5); z := 0; t := 0

	fmt.Println(funcName(10))
	fmt.Println(x)
	fmt.Println(nameReturn("David"))
	first, second := moreReturns("David", "Petro")
	fmt.Println("Name : ", first, second)
	fmt.Println(variadicFunc(2, 10, 44, 22))

	add := func() int {
		z += 2; return z
	}
	addOther := func() int {
		t += 2
		return t
	}
	fmt.Println(add())
	fmt.Println(add())

	fmt.Println(addOther())
	fmt.Println(addOther())


	fmt.Println(funcInsideFunc())

}

func funcName(a int) int {
	return a * a
}

func nameReturn(a string) (x string) {
	x = a
	return
}

func moreReturns(a string, b string) (string, string) {
	return a, b
}

func variadicFunc(x ...int) int {
	var res int
	for _, v := range x {
		res += v
	}
	return res
}

func funcInsideFunc() func() int {
	x := 10
	return func() int {
		return x * x
	}
}