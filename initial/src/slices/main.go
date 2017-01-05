package main

import "fmt"

func main() {

	slice := make([]int, 6)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice,1)
	slice = append(slice,10,20)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice2 := make([]int, 2,5)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice2 = append(slice2,10,2,5,40)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	for i:=0; i < 20; i++ {
		slice2 =append(slice2,i)
		fmt.Println("Len: ", len(slice2), " Cap:", cap(slice2))
	}


	fmt.Println("Slice", slice)
	slice3 := slice
	slice = append(slice,3,44,69)
	slice[0] = 11
	fmt.Println("Slice", slice)
	fmt.Println("Slice3", slice3)


	sliceString := []string{
		"Hello",
		"World",
		"Much",
		"Better",
		"Better 2",
	}

	fmt.Println(sliceString)
	fmt.Println(sliceString[0]) // Hello
	fmt.Println(sliceString[:2]) // Hello World
	fmt.Println(sliceString[1:2]) // World
	fmt.Println(sliceString[2:4]) // Much Better
	fmt.Println(sliceString[2:]) // Much Better Better 2



}
