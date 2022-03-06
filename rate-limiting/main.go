package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

func main() {
	total, max := 10, 3

	var wg sync.WaitGroup

	for i := 0; i < total; i++ {
		limit := max

		if i+max > total {
			limit = total - 1
		}

		wg.Add(limit)

		for j := 0; j < limit; j++ {
			go func(j int) {
				defer wg.Done()

				conn, err := net.Dial("tcp", ":8080")

				if err != nil {
					log.Println(err)
				}

				b, err := ioutil.ReadAll(conn)
				if err != nil {
					log.Println(err)
				}

				if string(b) != "success" {
					log.Fatalf("requst error, request: %d", i+1+j)
				}
				fmt.Println("request", i+j+1)
			}(j)
		}

		wg.Wait()
	}
}
