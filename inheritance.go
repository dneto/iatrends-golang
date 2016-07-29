package main

import (
	"fmt"
)

type Person struct {
	Name string
	age  int
}

type Employee struct {
	Person //Campo "anônimo" (não possui uma variável associada) no struct
	Id     int
}

func main() {
	employee := Employee{Person{"John", 20}, 1}
	fmt.Println("employee.Name:", employee.Name)
	fmt.Println("employee.age:", employee.age)
	fmt.Println("employee.Id:", employee.Id)
}
