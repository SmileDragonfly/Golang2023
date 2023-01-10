package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}
	//colors := make(map[string]string)
	//colors["red"] = "#ff0000"
	//delete(colors, "red")
	//fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for color", key, "is", value)
	}
}
