package main

import (
	"fmt"
	"strconv"
)

var (
    a byte =  1
    b int  =  -11
    c uint =  22
    d bool = false
    e float32 = 1.24
    f int  = 65
)

func main() {
	g := strconv.Itoa(f)
    fmt.Println(g)
}