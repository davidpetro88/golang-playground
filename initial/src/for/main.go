package main

import "fmt"

func main() {

	for i :=0; i < 5; i++  {
		fmt.Println(i)
	}
	
	x := 0

	for x < 10 {
		fmt.Println(x)
		x++
	}

	a := 0
	for  {

		fmt.Println("Infinited for =%s",a)
		a++
		if a == 50 {
			fmt.Println("Continue Infinited for =%s",a)
			continue
		}


		if a == 100 {
			fmt.Println("Break Infinited for =%s",a)
			break
		}

	}


}