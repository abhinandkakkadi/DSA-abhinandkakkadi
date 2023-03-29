package main

import "fmt"

func main() {

	arr := []int {3,1,5,4,2}

	bubbleSort(arr)
	fmt.Println(arr)

	// bubbleSort Reverse order
	arr2 := []int{1,4,5,1,2}
	bubbleReverse(arr2)
	fmt.Println(arr2)

	// sort only a portion of the array specified from last
	bubblePortion(arr2,2)
	fmt.Println(arr2)

}

func bubbleSort(arr []int) {
	
	for i:=0; i< len(arr); i++ {

		swap := false
		for j:=1; j< len(arr) - i; j++ {
			
			if arr[j] < arr[j-1] {
				swap = true
				arr[j],arr[j-1] = arr[j-1],arr[j]
			}

		}
		if !swap {   // If swap is false that the array is sorted
			break
		} 
	}
}

// Bubble sort reverse

func bubbleReverse(arr2 []int) {

	for i:=0; i < len(arr2); i++ {

		swap := false

		for j:=1; j< len(arr2) - i; j++ {
			
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
func bubblePortion(arr2 []int,portion int) {

	for i:=0; i < portion; i++ {
		swap := false
		for j:=1; j < len(arr2) - i; j++ {
			
			if arr2[j] < arr2[j-1] {
				swap = true
				arr2[j], arr2[j-1] = arr2[j-1],arr2[j]
			}

		}

		if !swap {
			break
		}
	}
}
