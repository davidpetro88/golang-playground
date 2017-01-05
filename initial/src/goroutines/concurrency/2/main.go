package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	"runtime"
)

var waitGroup sync.WaitGroup

func init()  {
	//não é necessario a partir do go 1.5
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println(runtime.NumCPU())
	waitGroup.Add(2)
	go runProcess("P1",20)
	go runProcess("P2",20)
	waitGroup.Wait()
	var s string
	fmt.Scanln(&s)
}

func runProcess(name string, total int)  {
	for i :=0; i < total; i++{
		fmt.Println(name,"->",i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
	waitGroup.Done()
}

