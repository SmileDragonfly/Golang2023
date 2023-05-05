package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1)
	fmt.Println("Start main")
	go func() {
		fmt.Println("Start func")
		time.Sleep(1000 * time.Millisecond)
		c <- 1
		c <- 2
		fmt.Println("End func")
	}()
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("End main")
}
