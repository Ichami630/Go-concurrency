package main

import (
	"fmt"
	"sync"
)

// rate limiting concurrent request execution

type request func()

func main() {
	requests := map[int]request{}
	//creates a map where each key is an int and each value is a request(which is just a func)

	for i := 1; i <= 100; i++ {
		f := func(n int) request {
			return func() {
				fmt.Println("request", n)
			}
		}

		requests[i] = f(i)
	}
	//generate 100 mock requests, each printing it number,each one is stored in map

	var wg sync.WaitGroup //waitgroup that will help us wait for each batch of 10 request to finish
	max := 10             //set the rate limit per batch to 10

	//looping all 100 request in chunks of 10
	for i := 0; i < len(requests); i += max {
		for j := i; j < i+max; j++ {
			wg.Add(1)
			go func(r request) {
				defer wg.Done()
				r()
			}(requests[j+1])
		}
		//allows max number of go routines to run in a batch
		wg.Wait()
		fmt.Println(max, "Request processed")
	}
	//after lauching a batch of 10, the program waits for all to finish, then moves to the next 10 batch

}
