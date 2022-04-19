package main

import "fmt"

func main() {
	// arr := [6][6]int{
	// 	{-9, -9, -9, 1, 1, 1},
	// 	{0, -9, 0, 4, 3, 2},
	// 	{-9, -9, -9, 1, 2, 3},
	// 	{0, 0, 8, 6, 6, 0},
	// 	{0, 0, 0, -2, 0, 0},
	// 	{0, 0, 1, 2, 4, 0},
	// }

	arr := [][]int{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}

	var result int

	for i := 0; i < 4; i++ {
		if i == 0 {
			result = arr[0][0] + arr[0][1] + arr[0][2] +
				arr[1][1] +
				arr[2][0] + arr[2][1] + arr[2][2]
		}

		for j := 0; j < 4; j++ {
			t :=
				// line 1
				arr[i][j] + arr[i][j+1] + arr[i][j+2] +
					// line 2
					arr[i+1][j+1] +
					// line 3
					arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if t > result {
				result = t
			}
		}
	}

	fmt.Println(result)

}
