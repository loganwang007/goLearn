package main

import "fmt"

func main() {
	color := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffff",
	}

	// color := make(map[string]string)

	// color["white"] = "#ffff"

	// delete(color, "white")

	printMap(color)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
