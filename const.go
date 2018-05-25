package main

import (
	"fmt"
)

//const a int = 1

const (
	//在go中方法常量首字母大写外部就可引用 反之则不可以引用 但是常量通常为大写 这个时候我们不想让常量呗外部使用 我们在常量前加 _ 或者加 c 表示 const
	//常量中只能使用常量 和 内置函数
	_A = 1
	//iota 从0开始 每定义一个常量 iota+1 遇到const iota重置为0 下行的iota = 1
	// << 左移动多少位
	_B int = _A << (iota * 10)
	//这里的_C 等上行 变量表达式  === _C = _A >> (iota * 10)
	_C
	_D, _E = 1, 2
	//这里就不能使用上上行表达式了 因为上行表达式是并行赋值
	//_F
	//如果上行是并行赋值 本行也只能并行赋值
	_F, _G
)

func main() {
	fmt.Println(_A)
	fmt.Println(_B)
	fmt.Println(_C)
	fmt.Println(_D)
	fmt.Println(_E)
	fmt.Println(_F)
	fmt.Println(_G)
	fmt.Println(1024 / 2)

	fmt.Println(^1024)
	fmt.Println(!true)
	fmt.Println(2 ^ 10)

	//和运算
	fmt.Println(6 & 11)
	//或运算
	fmt.Println(6 | 11)
	//差运算
	fmt.Println(6 ^ 11)
	//和差运算
	fmt.Println(6 &^ 11)

}

//   1   10  100 1000  10
//1,048,576
//000000000
// 11   1100
