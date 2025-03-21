package main

import (
	"fmt"
	"time"
)

// asynchronous code execution example
func main() {
	now := time.Now() //get the current time
	go task1()
	go task2()
	go task3()
	go task4()

	// time.Sleep(time.Second) //sleep for 1s
	//now all of the tasks are printed but this is not still the best way to acheive
	//and can be solved using the "fork join model"
	fmt.Printf("%v elapsed", time.Since(now)) //this will take approx 17.012nano seconds to complete execution for much better than the sync which takes 401.5ms
	//now see that none of the above taskes are printed

	//now the common question should be why does it exit without waiting on the child process to finish?

	/*This is due to how concurrency is done in Go.
	Namely in Go concurrency is done in conformance with something called FORK JOIN model*/

	//this can also be resolved using wait group join points or channels join points

}

func task1() {
	time.Sleep(100 * time.Millisecond) //wait for 100ms before exectuing
	fmt.Println("task 1")
}
func task2() {
	time.Sleep(200 * time.Millisecond) //wait for 200ms before exectuing
	fmt.Println("task 2")
}
func task3() {
	fmt.Println("task 4") //executes immediately
}
func task4() {
	time.Sleep(100 * time.Millisecond) //wait for 100ms before exectuing
	fmt.Println("task 4")
}
