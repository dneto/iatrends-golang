package main

import "fmt"

type Device interface {
	Info() string
}

type TV struct {
	Vendor string
}

func (tv *TV) Info() string {
	return tv.Vendor
}

func main() {
	var tv Device = &TV{"LG"}
	fmt.Println("tv.Info(): ", tv.Info())
}
