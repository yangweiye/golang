//声明包名称
package main
//引入包文件
import (
	"fmt"
)
//定义常量
const (
	IP = "127.0.0.1"
	ID = 1
)
//定义变量
var (
	name = "yang"
	age = 20
)
//声明类型
type (
	newType int
	newName int
)
//定义结构体
type gopher struct{}
//定义接口
type golang interface{}
//声明一个方法
func main() {
	fmt.Println("hello world")	
}