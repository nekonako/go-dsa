package main

import "fmt"

func main() {
	arr := [9]int{23, 12, 43, 2, 5, 1, 67, 88, 50}

	i := 1
	tmp := 0
	for i < len(arr) {
		if arr[i] >= arr[i-1] {
			i++
		} else {
			tmp = arr[i]
			arr[i] = arr[i-1]
			arr[i-1] = tmp
			if i > 1 {
				i--
			}
		}
	}

	fmt.Println(arr)

}
