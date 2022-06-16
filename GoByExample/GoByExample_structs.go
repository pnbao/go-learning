package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 40
	return &p
}

func main() {
	fmt.Println(person{"Bao", 26})
	fmt.Println(person{name: "Bao2", age: 27})
	fmt.Println(person{name: "Bao3"})
	fmt.Println(&person{name: "Bao4", age: 15})
	fmt.Println(newPerson("Baonew"))

	s := person{name: "Bao5", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 20
	fmt.Println(sp.age)
}
