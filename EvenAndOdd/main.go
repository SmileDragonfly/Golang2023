package main

import "fmt"

func main() {
	sn := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, it := range sn {
		if it%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}
