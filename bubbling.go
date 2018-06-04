package main

import "fmt"

func main() {
	var arr = [...]int{6, 12, 8, 2, 6, 4, 1}
	var l = len(arr)
	for o := 1; o < l; o++ {
		for i := 0; i < l-o; i++ {
			if arr[i] > arr[i+1] {
				var s = arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = s
			}
		}
	}

	fmt.Println(arr)

}
