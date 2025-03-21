package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4) //launch 4 goroutines

	now := time.Now()

	for i := 1; i < 5; i++ {
		go work(&wg, i)
	}

	wg.Wait()
	fmt.Println("Time elapsed..", time.Since(now)) //the total it takes the goroutines to run
	fmt.Println("Done wiating, main func exits")

}

func work(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task ", id, " is done")
}
