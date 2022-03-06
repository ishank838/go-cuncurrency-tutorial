package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)

	//an example of closure
	go func() {
		defer wg.Done()
		task1()
	}()

	wg.Wait()
	fmt.Println("elspsed", time.Since(t))
	fmt.Println("done waiting")
}

func task1() {
	time.Sleep(time.Second * 3)
	fmt.Println("task 1")
}
