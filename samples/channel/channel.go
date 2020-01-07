package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	quit := make(chan bool)

	go func(msg chan string) {
		for i := 0; i < 100; i++ {
			msg <- fmt.Sprintf("Send %d", i)
			time.Sleep(2 * time.Millisecond)
		}
		quit <- true
	}(c)

	go func(msg chan string) {
		for true {
			fmt.Println("Rec:1 > ", <-msg)
		}
	}(c)

	go func(msg chan string) {
		for true {
			fmt.Println("Rec:2 > ", <-msg)
		}
	}(c)

	<-quit
}
