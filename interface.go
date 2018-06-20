package main

import "fmt"

//这是一个空接口 可以说 所有方法和类型都实现了这个接口
type empty interface {
}

//定义一个接口 一个 struct 实现了接口中所有的方法就等于实现了该接口
type USB interface {
	//方法名 返回值类型 不写则无返回
	Name() string
	CONNECTER
}

// 接口可以嵌套关系
type CONNECTER interface {
	Connecter()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connecter() {
	fmt.Println("链接上了", pc.name)
}

type ComputerConnecter struct {
	name string
}

func (pc ComputerConnecter) Name() string {
	return pc.name
}

func (pc ComputerConnecter) Connecter() {
	fmt.Println("链接上了", pc.name)
}

func main() {
	//这里明确指出 a 实现USB接口
	var a USB
	a = PhoneConnecter{
		name: "手机",
	}
	fmt.Println(a.Name())
	a.Connecter()

	//这里b实现了 USB接口的所有方法等于实现了 UBS接口
	b := ComputerConnecter{
		name: "电脑",
	}
	b.Connecter()
	check(a)
	check(b)
}

func check(usb USB) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println(pc.name, "type")
		return
	}
	fmt.Println(usb.Name(), "实现了USB接口")
}
