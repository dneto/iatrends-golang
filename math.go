package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func div(x, y int) (int, int) {
	if y == 0 {
		panic("Divisão por Zero")
	}

	return x / y, x % y
}

func main() {
	soma := add(10, 5) // soma => 15
	fmt.Println("Soma:", soma)

	q, r := div(10, 3) // q => 3, r => 1
	fmt.Println("Quociente:", q, "Resto:", r)
}
