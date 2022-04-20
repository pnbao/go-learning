package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key1"] = 7
	m["key2"] = 13

	fmt.Println("map:", m)

	value1 := m["key1"]
	fmt.Println("value1: ", value1)
	fmt.Println("len: ", len(m))

	delete(m, "key2")
	fmt.Println("map:", m)

	_, present := m["key2"]
	fmt.Println("present:", present)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
