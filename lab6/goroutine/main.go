package main

import (
	"log"
	"time"
)

func countdown() {
	for i := 5; i >= 0; i-- {
		log.Printf("%d...", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go countdown()
	go func() {
		time.Sleep(time.Second * 2)
		log.Println("Delayed Greetings!")
	}()
	time.Sleep(time.Second * 4)
}
