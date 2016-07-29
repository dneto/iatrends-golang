package main

import "fmt"

func main() {
	esseAno := 2016
	p := &esseAno
	fmt.Println(*p)
	*p = 2017
	fmt.Println(esseAno)
}
