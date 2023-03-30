package main

import "fmt"

func main() {

	// quick sort
	arr := []int{4,2,1,465,2,1}
	quicksort(arr,0,len(arr)-1)
	fmt.Println(arr)



	// quick sort reverse
	arr2 := []int{5,4,3,22,4,1}
	quickReverse(arr2,0,len(arr)-1)
	fmt.Println(arr2)


	// string sort

	str := "abhinand"
	str2 := []byte(str)
	stringQuick(str2,0,len(str2)-1)
	fmt.Println(string(str2))
}

func quicksort(arr []int, low int, high int) {

	if low < high {

		pivot := partition(arr,low,high)

		quicksort(arr,pivot+1,high)
		quicksort(arr,low,pivot-1)
	}
}

func partition(arr []int, low int, high int) int {

	pivot := arr[high]
	i := low -1

	for j :=low; j< high; j++ {

		if arr[j] < pivot {
			i++

			arr[i],arr[j] = arr[j],arr[i]
		}

	}

	arr[i+1],arr[high] = arr[high],arr[i+1]

	return i+1
}



// quick reverse
func quickReverse(arr2 []int, low int, high int) {

	if low < high {

		pivot := partitionReverse(arr2,low,high)
		quickReverse(arr2,pivot+1,high)
		quickReverse(arr2,low,pivot-1)


	}
}

func partitionReverse(arr2 []int, low int, high int) int {

	pivot := arr2[high]
	i := low - 1
  
	for j:= low; j < high; j++ {

		if arr2[j] > pivot {
			i++

			arr2[i],arr2[j] = arr2[j],arr2[i]
		}
	}

	arr2[i+1],arr2[high] = arr2[high],arr2[i+1]

	return i+1
}



// 
func stringQuick(str2 []byte, low int, high int) {

	if low < high {

		pivot := partitionString(str2,low,high)
		stringQuick(str2,pivot+1,high)
		stringQuick(str2,low,pivot-1)
	}

}

func partitionString(str2 []byte, low int, high int) int {

	pivot := str2[high]

	i := low -1

	for j := low; j < high; j++ {

		if str2[j] < pivot {

			i++

			str2[i],str2[j] = str2[j],str2[i]

		}

	}

	str2[i+1],str2[high] = str2[high], str2[i+1]

	return i + 1

}