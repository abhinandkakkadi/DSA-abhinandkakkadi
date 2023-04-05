package main

import "fmt"

func main() {

	arr := []int {34,234,1,123,2,23}

	mergeSort(arr,0,len(arr)-1)
	fmt.Println(arr)
}

func mergeSort(arr []int, low int, high int) {

	if low >= high {
		return
	}

	mid := (low +high)/2

	mergeSort(arr,low,mid)
	mergeSort(arr,mid+1,high)
	merge(arr,low,mid,high)

}

func merge(arr []int, low int, mid int, high int) {

	i := low
	j := mid + 1
	var b []int

	for i<=mid && j <=high {

		if arr[i] < arr[j] {
			b = append(b, arr[i])
			i++
		} else {
			b = append(b, arr[j])
			j++
		}
	}

	for ;i<=mid; i++ {
		b = append(b, arr[i])
	}

	for ;j<=high; j++ {
		b = append(b, arr[j])
	}

	for i:=low; i<=high; i++ {
		arr[i] = arr[i-low]
	}
}