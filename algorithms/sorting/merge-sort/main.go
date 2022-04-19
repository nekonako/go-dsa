package main

import "fmt"

func main() {
	arr := []int{23, 12, 43, 2, 5, 1, 67, 88, 50}

	sorted := mergeSort(arr)

	fmt.Println(sorted)

}

func merge(left []int, right []int) []int {
	counter := 0
	result := make([]int, 0, len(left)+len(right))
	fmt.Println("cokk")
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}

		if left[0] <= right[0] {
			counter += 1
			result = append(result, left[0])
			left = left[1:]
		} else {
			counter += 1
			result = append(result, right[0])
			right = right[1:]
		}

	}

	fmt.Println("cok", counter)
	return result

}

func mergeSort(arr []int) []int {
	counter := 0
	if len(arr) <= 1 {
		return arr
	}
	counter += 1
	middle := len(arr) / 2
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])
	fmt.Println(counter)
	return merge(left, right)

}
