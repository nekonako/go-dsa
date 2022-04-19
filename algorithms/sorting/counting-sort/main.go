package main

import "fmt"

// get largest number from a slice
func getLargestItm(arr [9]int) int {

	if len(arr) == 0 {
		return 1
	}

	l := arr[0]

	for _, v := range arr {
		if v > l {
			l = v
		}
	}

	return l + 1

}

func main() {
	arr := [9]int{23, 12, 43, 2, 5, 1, 67, 88, 50}

	l := getLargestItm(arr)

	countSort := make([]int, l)

	// change item to 1 if the item of arr is same with he index of countSort
	for i := 0; i < len(arr); i++ {
		countSort[arr[i]] += 1
	}

	for i, j := 0, 0; i < l; i++ {
		for {
			if countSort[i] > 0 {
				arr[j] = i
				j += 1
				countSort[i] -= 1
				continue
			}
			break
		}
	}

	fmt.Println(arr)

}
