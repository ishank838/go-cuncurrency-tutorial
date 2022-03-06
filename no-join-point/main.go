package main

import (
	"fmt"
	"time"
)

// no waitgroup & channels
func main() {
	//t := time.Now()

	go task1()

	time.Sleep(time.Second * 1)
	fmt.Println("done waiting")
	//fmt.Println("elspsed", time.Since(t))
}

func task1() {
	time.Sleep(time.Second * 3)
	fmt.Println("task 1")
}
