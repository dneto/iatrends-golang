package main

import "fmt"

func main() {
	fruitColor := map[string]string{
		"Apple":  "Red",
		"Banana": "Yellow",
	}

	fmt.Println("Apple color is: ", fruitColor["Apple"])

	fruitColor["Orange"] = "Orange"

	fmt.Println("Orange color is: ", fruitColor["Orange"])
}
