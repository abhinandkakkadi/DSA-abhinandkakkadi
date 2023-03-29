package main

import "fmt"

func main() {

	arr := []int{1,5,2,25,222,23}

	insertionSort(arr)
	fmt.Println(arr)

	insertionSortDescending(arr)
	fmt.Println(arr)

}

func insertionSort(arr []int) {

	var current int
	var j int
	for i:=1; i< len(arr); i++ {
		current = arr[i]
		j = i-1

		for j>=0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = current
	}
}

func insertionSortDescending(arr []int) {

	var current int
	var j int
	for i:=1; i< len(arr); i++ {

		current = arr[i]
		j = i-1

		for j >=0 &&  current > arr[j] {
			
			arr[j+1] = arr[j]
			j--

		}
		arr[j+1] = current
	}
}