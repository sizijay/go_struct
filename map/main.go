package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#FF0000	",
		"green": "#008000	",
		"white": "#000000	",
		"yellow": "#FFFF00	",
		"orange": "#FFA500	",
	}

	print(colors)
}

func print(c map[string]string) {
	for color, hex := range c {
		fmt.Println("The hex code for the color " + color + " is " + hex)
	}
}
