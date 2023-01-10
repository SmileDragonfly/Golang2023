package main

import "fmt"

func main() {
	//colors := map[string]string{
	//	"red":   "#ff0000",
	//	"green": "#00ff00",
	//}
	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	delete(colors, "red")
	fmt.Println(colors)
}
