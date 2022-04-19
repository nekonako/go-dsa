package main

import "fmt"

/// coktail sort seperti bubble sort
// bedanya array dibagi menjadi 2, lalu di urutkan dari ke 2 sisi

func main() {
	arr := [9]int{23, 12, 43, 2, 5, 1, 67, 88, 50}
	tmp := 0

	for i := 0; i <= len(arr)/2; i++ {
		left := 0
		right := len(arr) - 1

		for left <= right {
			// dari depan
			// jika array saat ini lebih dari array berikutnya
			if arr[left] > arr[left+1] {
				// simpan array saat ini ke temporary variabe
				tmp = arr[left]
				// pindahkan array berikutnya ke array saat ini
				arr[left] = arr[left+1]
				// pindahkan array saat ini yang disimpan di tmp ke array berikutnya
				arr[left+1] = tmp
			}

			left++

			// dari belakang
			// jika array berikutnya (dari belakang) lebih dari array saat ini
			if arr[right-1] > arr[right] {
				// simpan array berikunya ke tmp
				tmp = arr[right-1]
				// pndahkan array saat ini ke berikunya (dari belakang)
				arr[right-1] = arr[right]
				// ambil value mp dan pindahkan ke array saat ini
				arr[right] = tmp
			}

			right--

		}
	}

	fmt.Println(arr)

}
