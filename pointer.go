package main

import "fmt"

const ()

func main() {
	var a int = 123
	//递增递减需要单独写
	a++;
	//获取指针符号 & 存储指针类型 *int
	var p *int = &a
	//获取指针存储内容 *指针
	if *p == a {
		var b = 1
		fmt.Println(*p)
		fmt.Println(b)
	}

	//if判断 可先初始化参数 参数局部（if内）有效
	if c := 111; c >= 1 {
		fmt.Println(c)
	}

	//for循环 条件循环
	for a > 1 {
		fmt.Println("条件循环")
		a = a - 30
	}

	//经典for循环
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		fmt.Println("经典循环")
	}

	var b int = 10
	//经典switch 判断
	switch b {
	case 1:
		fmt.Println("b = 1")
	case 10:
		fmt.Println("b = 10")
	default:
		fmt.Println("b undefined")
	}

	//条件switch 判断
	switch {
	case b > 1:
		fmt.Println("b = 1")
		//继续执行下一个符合条件的分支 包括default
		fallthrough
	case b > 2:
		fmt.Println("b = 1")
		fallthrough

	case b > 3:
		fmt.Println("b = 1")
		fallthrough

	case b < 30:
		fmt.Println("b = 1")
		fallthrough

	default:
		fmt.Println(11)
	}

	//标签
Label1:
	for {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			//跳出循环 到指定标签
			break Label1
			//将代码运行从 Label2 标签继续执行
			goto Label2
		}
	}
Label2:

	for v := 0; v < 10; v++ {
		for {
			fmt.Println(v)
			//跳出指定循环到 标签
			continue Label2
		}
	}

}
