package main

import "fmt"

func main() {
	arr := [10]int{100, 78, 12, 34, 55, 76, 88, 90, 22, 21}

	// loop array
	for i := 0; i < len(arr); i++ {
		// loop each item on array and compare wih next value
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)

}
