package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	now := time.Now()

	wg.Add(1) //we are launching 1 go routine
	go func() {
		defer wg.Done()
		work()
	}()

	wg.Wait() //wait for all goroutines to finish
	fmt.Println("time waited...", time.Since(now))
	fmt.Println("the main go routine")

}

func work() {
	time.Sleep(500 * time.Millisecond) //pause for 500ms
	fmt.Println("work function")
}
