package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	fmt.Println("Start main")
	go func() {
		fmt.Println("Start func")
		time.Sleep(10000 * time.Millisecond)
		c <- 1
		fmt.Println("End func")
	}()
	<-c
	fmt.Println("End main")
	fmt.Println("Channel value: ", c)
}
