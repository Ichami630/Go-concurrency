package main

import (
	"fmt"
	"time"
)

// synchronous code execution example
func main() {
	now := time.Now() //get the current time
	task1()
	task2()
	task3()
	task4()

	fmt.Printf("%v elapsed", time.Since(now)) //this will take approx 401.5ms to complete execution

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
