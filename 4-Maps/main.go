package main

import "fmt"

func main() {
	colors := map[string]string{ //First way to declare a map
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//var colors map[string]string //Second way to declare a map
	//colors := make(map[int]string) //Third way to declare a map

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		//fmt.Printf("Color: %s, Hex: %s\n", color, hex)
		fmt.Printf("Color: %s, Hex: %s\n", color, hex)
	}
}
