package array


func LargestElement(arr []int) int {
	largest := arr[0]
	for i := 1 ; i < len(arr); i++ {
		if largest < arr[i] {
			largest = arr[i]
		}
	}
	return largest
}

func SecondLargest(arr []int) int {

	var largest int
	var secondLargest int

	if arr[0] > arr[1] {
		largest = arr[0]
		secondLargest = arr[1]
	} else {
		largest = arr[1]
		secondLargest = arr[0]
	}

	for i := 2; i < len(arr); i++ {
		
		if arr[i] > largest {

			secondLargest = largest
			largest = arr[i]

		} else if arr[i] > secondLargest {

			secondLargest = arr[i]
			
		}
	}

	return secondLargest

}

// Remove duplicates from sorted array

func RemoveDuplicatesFromSortedArray(arr []int) int {

		temp := arr[0]
		j := 0

		for i := 1 ; i < len(arr); i++ {

			if arr[i] != temp {
				j++
				arr[j] = arr[i]
				temp = arr[j]
			}
		}

		return j+1
}

// left rotate array

func LeftRotate(arr []int, n int) {
	
	if n > len(arr)-1 {
		n = n%len(arr)-1
	}

	for n > 0 {

		temp := arr[0]
		for i :=0; i < len(arr)-1 ; i++ {
			arr[i] = arr[i+1]
		}	
		n--
		arr[len(arr)-1] = temp

	}
}

// Move zeroes to end

func MoveZeroes(arr []int) {

	j := -1

	for i := 0; i < len(arr); i++ {

		if arr[i] != 0 {
			j++
			arr[j],arr[i] = arr[i],arr[j]
		}
		
	}
}

// union of two array
func UnionArray(arr1 []int, arr2 []int) []int {

	l1 := len(arr1)
	l2 := len(arr2)
	var arr3 []int
	k := 0
	i := 0
	j := 0
	for i < l1 && j < l2  {

		if arr1[i] <= arr2[j] {
			if k==0 || arr3[k-1]!= arr1[i] {
				arr3 = append(arr3, arr1[i])
				k++
				
			}
			i++
			
		} else {
			if k==0 || arr3[k-1]!= arr2[j] {
				arr3 = append(arr3, arr2[j])
				k++
				
			}
			j++
		}

		

	}

	for ;i < l1; i++ {
		if k==0 || arr3[k-1]!= arr1[i] {
			arr3 = append(arr3, arr1[i])
			k++
		}
	}

	for ;j < l2; j++ {
		if k==0 || arr3[k-1]!= arr2[j] {
			arr3 = append(arr3, arr2[j])
			k++
		}
	}

	return arr3

}


