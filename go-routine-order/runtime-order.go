package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go task1(&wg)
	wg.Wait()

	wg.Add(2)
	go task2(&wg)
	go task3(&wg)
	wg.Wait()

	wg.Add(1)
	go task4(&wg)
	wg.Wait()

	wg.Add(1)
	go task5(&wg)
	wg.Wait()
}

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Println("task1 done")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Println("task2 done")
}
func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Println("task3 done")
}
func task4(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Println("task4 done")
}
func task5(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Println("task5 done")
}
func task6(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Println("task6 done")
}

func task7(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Println("task7 done")
}
