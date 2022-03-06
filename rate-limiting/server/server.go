package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}
	shutdown(l)

	connections := int32(0)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		connections++

		go func() {
			//serve connection

			defer func() {
				conn.Close()
				atomic.AddInt32(&connections, -1)
			}()

			if atomic.LoadInt32(&connections) > 3 {
				return
			}

			time.Sleep(time.Second * 5)
			_, err := conn.Write([]byte("success"))
			if err != nil {
				log.Println(err)
			}
		}()
	}
}

func shutdown(l net.Listener) {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	go func() {
		for range ch {
			l.Close()
			<-ch
		}
	}()

}
