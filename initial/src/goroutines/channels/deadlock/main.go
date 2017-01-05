package main

import "fmt"

func main() {
	channel := make(chan int)
	//channel <- 10 // Vai dar o erro de deadlock, pois o channel teria que estar dentro de um go routing
	go func() {
		channel <- 10
	}()

	fmt.Println(<-channel)
}

