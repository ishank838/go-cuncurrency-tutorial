package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	done := make(chan struct{})

	//an example of closure
	go func() {
		task1()
		done <- struct{}{}
	}()

	<-done
	fmt.Println("elspsed", time.Since(t))
	fmt.Println("done waiting")
}

func task1() {
	time.Sleep(time.Second * 3)
	fmt.Println("task 1")
}
