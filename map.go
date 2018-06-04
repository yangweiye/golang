package main

import (
	"fmt"
	"sort"
)

func main() {
	var m map[int]string
	fmt.Println(m)
	m = make(map[int]string)
	fmt.Println(m)
	m[1] = "map"
	fmt.Println(m)
	a := m[1]
	fmt.Println(a)
	delete(m, 1)
	fmt.Println(m)

	var b map[int]map[int]map[int]string
	b = make(map[int]map[int]map[int]string)
	b[1] = make(map[int]map[int]string)
	b[1][1] = make(map[int]string)
	b[1][1][1] = "1ok"
	fmt.Println(b)
	c, ok := b[1][1][1]
	fmt.Println(c, ok)
	c, ok = b[2][1][1]
	if !ok {
		b[2] = make(map[int]map[int]string)
		b[2][1] = make(map[int]string)
		b[2][1][1] = "2ok"
	}
	c, ok = b[2][1][1]
	fmt.Println(c, ok)

	for k, v := range b {
		fmt.Println(k)
		fmt.Println(v)
	}

	var sm = make([]map[int]string, 5)

	for i := range sm {
		sm[i] = make(map[int]string)
		sm[i][1] = "ok"
	}

	fmt.Println(sm)

	var d = map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	fmt.Println(d)
	var s = make([]int, len(d))
	var index int
	for key := range d {
		s[index] = key
		index++
	}
	sort.Ints(s)
	fmt.Println(s)
	fmt.Println(d)

	var e = make(map[string]int)
	for k1, v1 := range d {
		e[v1] = k1
	}

	fmt.Println(e)
}
