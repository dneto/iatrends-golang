package main

import "fmt"

type Person struct {
	Name string
	age  int
}

type Employee struct {
	Person
	employeeId int
}

func main() {
	person := Person{Name: "Methuselah", age: 969}
	personPointer := &Person{Name: "John", age: 20}

	name, age := person.Name, person.age
	pointerName, pointerAge := personPointer.Name, personPointer.age

	fmt.Println("name:", name, ", age:", age)
	fmt.Println("pointerName:", pointerName, ", pointerAge:", pointerAge)
}
