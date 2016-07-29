package main

import "fmt"

type IntList []int

func (iL IntList) first() int {
	return iL[0]
}

func m() {
	var list IntList = []int{1, 2, 3, 4, 5}
	fmt.Println(list.first())
}
