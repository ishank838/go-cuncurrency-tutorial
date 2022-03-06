package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type WaitGroup struct {
	count int64
}

func (wg *WaitGroup) Add(inc int64) {

	atomic.AddInt64(&wg.count, inc)
}

func (wg *WaitGroup) Done() {
	if atomic.LoadInt64(&wg.count) < 0 {
		panic("negative waitgroup counter")
	}
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	for {
		if atomic.LoadInt64(&wg.count) == 0 {
			return
		}
	}
}

func main() {
	var wg WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 5)
		fmt.Println("done 1")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		fmt.Println("done 2")
	}()
	wg.Wait()
}
