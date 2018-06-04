package main

import "fmt"

func main() {
	s0 := []int{}
	fmt.Println(s0)
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := a[0]
	fmt.Println(s1)
	s2 := a[0:5]
	fmt.Println(s2)
	s3 := a[:5]
	fmt.Println(s3)
	s4 := a[5:10]
	fmt.Println(s4)
	s5 := a[5:]
	fmt.Println(s5)

	//创建一个 长度为3 容量为 10 的切片（slice） 当长度超过容量时 容量会翻倍
	ss := make([]int, 3, 10)
	fmt.Println(ss)
	fmt.Println(len(ss), cap(ss))

	ss1 := a[2:3]
	//索引不能大于其容量  否则会报错而不是从新分配内存
	//fmt.Println(ss1[1])
	//slice 容量等于 被切片的 cap
	fmt.Println(ss1, len(ss1), cap(ss1))
	fmt.Printf("%p \n", ss1)

	//长度不超过容量就不会从新分配内存地址
	ss2 := append(ss1, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(ss2, len(ss2), cap(ss2))
	fmt.Printf("%p \n", ss2)

	//切片都是引用赋值 一个变另一个也变
	ss1[0] = 10
	fmt.Println(ss1, len(ss1), cap(ss1))
	fmt.Println(ss2, len(ss2), cap(ss2))

	//当切片长度超过容量 内存地址发送变化 就和原来的切片没有关系了
	ss2 = append(ss2, 1, 2, 3, 4, 5)
	ss1[0] = 20
	fmt.Println(ss1, len(ss1), cap(ss1))
	fmt.Printf("%p \n", ss1)
	fmt.Println(ss2, len(ss2), cap(ss2))
	fmt.Printf("%p \n", ss2)

	//copy 的使用
	ss3 := []int{1, 2, 3}
	fmt.Println(ss3, len(ss3), cap(ss3))
	ss4 := []int{4, 5, 6, 7, 8, 9}

	fmt.Println(ss4, len(ss4), cap(ss4))
	//将 ss4 copy 到 ss3 因为ss3 长度为3 所以只会赋值前3个
	copy(ss3, ss4)
	fmt.Println(ss3)
	ss3 = []int{1, 2, 3}
	//将 ss3 copy 到 ss4 中 因为ss3 只有3个值所以只会替换3个值 其他的不变
	copy(ss4, ss3)
	fmt.Println(ss4)
	ss4 = []int{4, 5, 6, 7, 8, 9}
	//指定ss3的哪几个值  赋值 到 ss4 的什么地方
	copy(ss4[3:5], ss3[1:3])
	fmt.Println(ss4)

	ss5 := []int{1, 2, 3, 4, 5}
	ss6 := ss5[0:]
	fmt.Println(ss5)
	fmt.Println(ss6)

}
