package main

import "fmt"

func main() {

	arr := []int{4,3,5,6,1,2}

	quickSort(arr,0,len(arr)-1)

	fmt.Println(arr)
}

func quickSort(arr []int,low int, high int) {

	if low < high {

		pivot := partition(arr,low,high)

		quickSort(arr,low,pivot-1)
		quickSort(arr,pivot+1,high)
	}
}

func partition(arr []int, low int, high int) int {

	pivot := arr[high]
	i := low -1

	for j := low; j < high; j++ {

		if arr[j] < pivot {
			i++

			arr[i],arr[j] = arr[j],arr[i]
		}
	}

	arr[i+1],arr[high] = arr[high],arr[i+1]

	return i + 1

}