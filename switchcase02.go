package main

import "fmt"

func main() {
	number := 10

	switch {
	case number < 5:
		fmt.Println("Number", number, "< 5")
	case number >= 6:
		fmt.Println("Number", number, ">=6")
	}
}
