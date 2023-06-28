package main

import "fmt"

func main() {

	arr := []int{5, 3, 756, 11, 5, 1}

	// searching
	selectionSort(arr)
	fmt.Println(arr)

	//
}

func selectionSort(arr []int) {

	var k, pos int

	for i := 0; i < len(arr); i++ {

		k = arr[i]
		pos = i

		for j := i + 1; j < len(arr); j++ {

			if arr[j] < k {

				k = arr[j]
				pos = j

			}

		}

		arr[i], arr[pos] = arr[pos], arr[i]
	}
}
