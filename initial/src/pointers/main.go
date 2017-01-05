package main

import "fmt"

func test(a int) int {
	return a + 1
}

func test2(a *int) int {
	*a = *a + 1
	return *a
}

func main() {
	x := 10
	fmt.Println(&x) // 0xc42000a360
	fmt.Println(x)

	y := &x
	fmt.Println(y)
	fmt.Println(*y)

	*y = 20
	fmt.Println(x)

	var z *int = &x
	fmt.Println(z)
	fmt.Println(*z)


	b :=10
	fmt.Println(test(b))
	fmt.Println(b)

	c :=10
	fmt.Println(test2(&c))
	fmt.Println(c)


}
