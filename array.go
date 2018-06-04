package main

import "fmt"

func main() {
	//基础定义 长度为2存储int的数字
	var a [2]int
	fmt.Println(a)

	//定义数组并赋值
	b := [2]int{1, 2}
	fmt.Println(b)

	//定义自动检测长度的数组并赋值 长度会自动计算
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	//定义自动检测长度的数组 并把它的第11位赋值为1
	d := [...]int{10: 1}
	fmt.Println(d)

	e := [2]int{1, 2}
	//数组key进行 == ！= 运算
	fmt.Println(b == e)

	//指向数组的指针
	f := &[...]int{1, 2, 3, 4, 5}
	fmt.Println(f)

	//多维数组 [2]表示一维数组里有几个数组 [3]表示每个二维数组中有几个值
	g := [2][3]int{
		{1, 1},
		{2, 2, 2}}
	g[0][1] = 3;
	fmt.Println(g)

	h := 1
	i := 2
	//指针行数组
	z := [2]*int{&h, &i}
	fmt.Println(z)

	p := new([10]int)
	//数组赋值
	p[1] = 2
	fmt.Println(p)

}
