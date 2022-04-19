package main

import "fmt"

func main() {
	arr := [9]int{23, 12, 43, 2, 5, 1, 67, 88, 50}
	tmp := 0

	arrLen := len(arr)
	gap := arrLen

	for gap > 1 {
		gap = gap * 10 / 13
		fmt.Println(gap)
		for i := 0; i+gap < arrLen; i++ {
			if arr[i] > arr[i+gap] {
				tmp = arr[i]
				arr[i] = arr[i+gap]
				arr[gap+i] = tmp
			}
		}
	}

	fmt.Println(arr)

}
