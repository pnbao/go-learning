package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

type describer interface {
	describe() string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "bao bao",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)
	fmt.Println("container:", co)
	fmt.Println("descibe:", co.describe(), co.base.describe())

	var d describer
	d = co
	fmt.Println("desciber:", d, d.describe())
}
