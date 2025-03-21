package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{}) //make a channel

	go func() {
		work()
		done <- struct{}{}
	}()

	<-done //wait for all go routines(forks) to complete

	fmt.Println("time elapsed..", time.Since(now))
	fmt.Println("done waiting, main func exits")

}

func work() {
	time.Sleep(500 * time.Millisecond) //pause for 500ms
	fmt.Println("work function")
}
