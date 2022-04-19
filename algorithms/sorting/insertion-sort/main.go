package main

import "fmt"

func main() {
	arr := [9]int{23, 12, 43, 2, 5, 1, 67, 88, 50}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println(arr)

}
