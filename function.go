package main

import "fmt"

func main() {
	fmt.Println(A(1, "2"))
	fmt.Println(B(1, 2, 3, 4))

	a := []int{1, 2, 3}
	b := 1
	C(a, &b)
	fmt.Println(a)
	fmt.Println(b)

	c := D
	c()

	d := func() {
		fmt.Println("this is a func d")
	}
	d()

	e := E(3)
	fmt.Println(e(3))

	F()
	G()
	H()

}

func A(a int, b string) (int, string) {
	a = 1
	b = "2"
	return a, b
}

func B(a ... int) ([]int) {
	return a
}

func C(a []int, b *int) (bool) {
	a[0] = 333
	*b = 111

	return true
}

func D() {
	fmt.Println("this is function D")
}

func E(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func F() {
	fmt.Println("this is function F")
}

func G() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in function G")
		}
	}()
	fmt.Println("this is function G")
	panic("panic in function G")
}
func H() {
	fmt.Println("this is function H")
}
