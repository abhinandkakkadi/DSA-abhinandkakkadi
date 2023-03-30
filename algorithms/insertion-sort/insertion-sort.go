package main

import "fmt"

func main() {

	arr := []int{1,5,2,25,222,23}

	// Ascending order
	insertionSort(arr)
	fmt.Println(arr)

	// Descending order
	insertionSortDescending(arr)
	fmt.Println(arr)

	// sort string
	str := "athira"
	str2 := []byte(str)
	insertionString(str2)
	fmt.Println(string(str2))

}


// Ascending order

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

// descending order
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

// sort string
func insertionString(str2 []byte) {

	for i:=1; i < len(str2); i++ {

		current := str2[i]
		j := i -1

		for j >=0 && current < str2[j] {

			str2[j+1] = str2[j]
			j--

		}

		str2[j+1] = current


	}
}