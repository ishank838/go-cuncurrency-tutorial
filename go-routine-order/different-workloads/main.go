package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(5)

	go task1(&wg)
	go task2(&wg)
	go task3(&wg)
	go task4(&wg)
	go task5(&wg)

	wg.Wait()

	fmt.Println("Done")
}

func task1(wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := http.Get("localhost:8080/")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("task1 done")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()

	var count int

	for i := 0; i < 10000000; i++ {
		count += i
	}

	fmt.Println("task2 done")
}
func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Println("task3 done")
}
func task4(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 600)
	fmt.Println("task4 done")
}
func task5(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Println("task5 done")
}
