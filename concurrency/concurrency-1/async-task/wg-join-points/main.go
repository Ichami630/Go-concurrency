package main

import (
	"fmt"
	"time"
)

func main() {
	go work()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("the main go routine")

	//no joint point, so the the child process (work function) will
	//not be executed

}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("work function")
}
