package main

import "fmt"

func main() {

	arr := []int{3, 1, 5, 4, 2}

	bubbleSort(arr)
	fmt.Println(arr)

	// bubbleSort Reverse order
	arr2 := []int{1, 4, 5, 1, 2}
	bubbleReverse(arr2)
	fmt.Println(arr2)

	// sort only a portion of the array specified from last
	bubblePortion(arr2, 2)
	fmt.Println(arr2)

	// bubbleSort on strings
	str := "abhinand"
	str2 := []byte(str)
	bubbleStrings(str2)
	fmt.Println(string(str2))

	// Number of swaps required to sort the array
	arr3 := []int{1, 34, 2, 2, 564, 2}
	count := bubbleCount(arr3, 0)
	fmt.Println("Total swap done in this array = ", count, arr3)

}

// bubble sort
func bubbleSort(arr []int) {

	for i := 0; i < len(arr); i++ {

		swap := false
		for j := 1; j < len(arr)-i; j++ {

			if arr[j] < arr[j-1] {
				swap = true
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}

		}
		if !swap { // If swap is false that the array is sorted
			break
		}
	}
}

// Bubble sort reverse

func bubbleReverse(arr2 []int) {

	for i := 0; i < len(arr2); i++ {

		swap := false

		for j := 1; j < len(arr2)-i; j++ {

			if arr2[j] > arr2[j-1] {
				swap = true
				arr2[j], arr2[j-1] = arr2[j-1], arr2[j]
			}
		}

		if !swap {
			break
		}

	}
}

// sort only a portion of the array specified from last
func bubblePortion(arr2 []int, portion int) {

	for i := 0; i < portion; i++ {
		swap := false
		for j := 1; j < len(arr2)-i; j++ {

			if arr2[j] < arr2[j-1] {
				swap = true
				arr2[j], arr2[j-1] = arr2[j-1], arr2[j]
			}

		}

		if !swap {
			break
		}
	}
}

// bubble sort on strings

func bubbleStrings(str2 []byte) {

	var swap bool

	for i := 0; i < len(str2); i++ {
		swap = false
		for j := 0; j < len(str2)-i-1; j++ {

			if str2[j] > str2[j+1] {
				swap = true
				str2[j], str2[j+1] = str2[j+1], str2[j]

			}
		}

		if !swap {
			break
		}
	}
}

// Number of swaps required to sort an array

func bubbleCount(arr3 []int, n int) int {

	var swap bool

	for i := 0; i < len(arr3); i++ {

		swap = false

		for j := 0; j < len(arr3)-i-1; j++ {

			if arr3[j] > arr3[j+1] {
				n++
				swap = true
				arr3[j], arr3[j+1] = arr3[j+1], arr3[j]
			}
		}

		if !swap {

			break

		}
	}

	return n
}
