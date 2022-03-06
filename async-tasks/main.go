package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	go task1()
	go task2()
	go task3()
	go task4()

	//does not print task since main goroutine escapes so
	//put some sleep
	time.Sleep(time.Second * 5)
	fmt.Println("elapsed", time.Since(t))

}

func task1() {
	time.Sleep(time.Second * 1)
	fmt.Println("task 1")
}

func task2() {
	time.Sleep(time.Second * 2)
	fmt.Println("task 2")
}

func task3() {
	fmt.Println("task 3")
}

func task4() {
	time.Sleep(time.Second * 4)
	fmt.Println("task 4")
}
