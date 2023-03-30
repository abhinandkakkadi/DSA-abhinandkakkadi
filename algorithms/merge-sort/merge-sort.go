package main

import (
	"fmt"
)


func main() {

	arr := []int{6,5,4,3,2,1,45,21,1,12}

	mergerSort(arr,0,len(arr)-1)
	fmt.Println(arr)
}

func mergerSort(arr []int,s int,l int) {

	if s >= l {
		return 
	}

	mid := (s+l)/2

	mergerSort(arr,s,mid)
	mergerSort(arr,mid+1,l)
	merge(arr,s,mid,l)
}

func merge(arr []int, s int, mid int, l int) {

	var b []int
	i:=s
	j:=mid+1

	for i <=mid && j <=l {

		if arr[i] <= arr[j] {
			b = append(b, arr[i])
			i++
		
		} else {
			b = append(b, arr[j])
			j++
		} 
	}

	for ; i<=mid; i++ {
		b = append(b, arr[i])
	}

	for ; j<=l; j++ {
		b = append(b, arr[j])
	}

	for i:=s; i<=l; i++ {
		arr[i] = b[i-s]
	}


}