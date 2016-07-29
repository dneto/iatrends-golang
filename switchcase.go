package main

import "fmt"

var fruitColor = map[string]string{
	"Apple":  "Red",
	"Banana": "Yellow",
	"Orange": "",
}

func main() {
	printFruitColor("A")
}

func printFruitColor(fruit string) {
	switch color := fruitColor[fruit]; color { // Atribuindo e utilizando na mesma linha
	case "Red":
		fmt.Println("Red")
	case "Yellow":
		fmt.Println("Yellow")
	default:
		fmt.Println("No color defined")
	}
}
