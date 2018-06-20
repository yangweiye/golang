package main

import "fmt"

type Aa struct {
	Name string
}

type num int

func main() {
	a := Aa{
		Name: "ywy",
	}
	a.Print()

	var b num = 1
	fmt.Println(b)
	b.Increease()
	fmt.Println(b)
}

func (b Aa) Print() {
	fmt.Println(b.Name)
}

func (number *num) Increease() {
	*number = *number + 100
}
