package main

import "fmt"

type person struct {
	name string
	age  int
}

type dog struct {
	color, volume string
	son struct {
		name string
		age  int
	}
}

type cat struct {
	name string
	age  int
}

type cat_king struct {
	cat
	name  string
	color string
}

type cat_god struct {
	cat
	level int
}

func main() {
	a := person{
		name: "yang weiye",
		age:  19,
	}
	//a.name = "Yang weiye"
	//a.age = 20
	fmt.Println(a)
	Aa(a)
	fmt.Println(a)

	b := &person{
		name: "liu",
		age:  10,
	}
	fmt.Println(b)
	Ab(b)
	fmt.Println(b)

	c := struct {
		name string
		age  int
	}{
		name: "dog",
		age:  2,
	}

	fmt.Println(c)

	big_dog := dog{
		color:  "红色",
		volume: "很大",
	}
	big_dog.son.name = "小狗"
	big_dog.son.age = 2

	fmt.Println(big_dog)

	small_cat := cat{
		"小猫",
		2,
	}

	fmt.Println(small_cat)

	mini_cat := cat{
		"迷你猫",
		1,
	}
	fmt.Println(small_cat == mini_cat)

	king := cat_king{
		color: "红色",
		name:  "猫王啊",
		cat:   cat{name: "猫王", age: 100},
	}
	god := cat_god{
		level: 1000,
		cat:   cat{name: "猫神", age: 10000},
	}
	fmt.Println(king)
	fmt.Println(god)

}

func Aa(per person) {
	per.age = 21
	fmt.Println(per)
}

func Ab(per *person) {
	per.age = 120
	fmt.Println(per)
}
