package main

import (
	"log"
	"time"
)

func startConsumer(queue chan int, name string) {
	for {
		time.Sleep(time.Second)
		log.Println(name, <-queue)
	}
}

func main() {
	queue := make(chan int, 10000)
	n := 10000

	for i := 1; i <= n; i++ {
		queue <- i
	}

	go startConsumer(queue, "C1")
	go startConsumer(queue, "C2")

	time.Sleep(time.Second * 100)
}
