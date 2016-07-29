package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Banana", "Orange"}
	fmt.Println("Fruit array", fruits)
	fruitSlice := fruits[:2]
	fruitSlice[0] = "Tomato"
	fmt.Println("Fruit slices:", fruitSlice)
	fmt.Println("Fruit array (updated):", fruits)
}
