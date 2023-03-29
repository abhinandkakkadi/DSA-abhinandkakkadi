package main

import "fmt"

func main() {

	arr := []int {1,2,3,4,5}

	bubbleSort(arr)
	fmt.Println(arr)
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