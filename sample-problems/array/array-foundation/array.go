package main

import (
	"fmt"
	"math"
)

func main() {

	// second smallest element in an array
	arr := []int{23, 1, 34, 34, 4}
	var smallest int = math.MaxInt
	var sSmallest int = math.MaxInt

	s, ss := second_smallest(arr, smallest, sSmallest)
	fmt.Println("The smallest and second smallest are ", s, " ", ss)

	// Reverse an array

	arr2 := []int{1, 2, 3, 4, 5, 6}
	reverse(arr2)
	fmt.Println(arr2)

	// recursive way to reverse an array

	arr3 := []int{1, 2, 3, 4, 5, 6}
	arr4 := recursiveReverse(arr3, 0, len(arr3)-1)
	fmt.Println("Recursive reverse ", arr4)

	// frequency of each element in an array

	arr5 := []int{1, 23, 12, 1, 223, 1, 23}
	frequency(arr5)

	// frequency using map

	arr6 := []int{4, 23, 1, 1, 34, 5, 4}
	hashFrequency(arr6)

	// Rearrange increasing-decreasing order
	arr7 := []int{4, 2, 8, 6, 15, 5, 9, 20}
	increasingDecreasing(arr7)
	fmt.Println(arr7)

	// Rotate array by k elements
	arr8 := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(arr8, k)
	fmt.Println(arr8)

	// median of an array
	arr9 := []int{1, 2, 3, 4, 5}
	m := median(arr9)
	fmt.Println("hey ", m)

	// remove duplicates from a sorted array
	arr10 := []int{1, 1, 2, 2, 3, 3, 4, 4}
	d, size := removeDuplicates(arr10, len(arr10))
	for i := 0; i < size; i++ {
		fmt.Println(d[i])
	}

	// quick sort
	arr11 := []int{33, 123, 1, 32, 112, 123, 1}
	quickSort(arr11, 0, len(arr)-1)
	fmt.Println(arr11)

	// GCD of 2 numbers
	num1, num2 := 3, 6
	fmt.Println(gcd(num1, num2))

	// Armstrong number
	num := 407
	fmt.Println(armStrong(num))
}

// second smallest element in an array
func second_smallest(arr []int, smallest int, sSmallest int) (int, int) {

	smallest = arr[0]

	for i := 1; i < len(arr); i++ {

		if arr[i] < smallest {
			sSmallest = smallest
			smallest = arr[i]
		} else if arr[i] < sSmallest && arr[i] != smallest {
			sSmallest = arr[i]
		}
	}

	return smallest, sSmallest
}

// Reverse an array

func reverse(arr2 []int) {

	for i, j := 0, len(arr2)-1; i < j; i, j = i+1, j-1 {
		arr2[i], arr2[j] = arr2[j], arr2[i]
	}
}

// Recursive reverse

func recursiveReverse(arr3 []int, s int, l int) []int {

	if s > l {

		return arr3

	}

	arr3[s], arr3[l] = arr3[l], arr3[s]
	return recursiveReverse(arr3, s+1, l-1)

}

// frequency

func frequency(arr []int) {

	brr := make([]int, len(arr))
	var count int
	for i := 0; i < len(arr); i++ {
		count = 1
		for j := i + 1; j < len(arr); j++ {

			if arr[i] == arr[j] && brr[j] != -1 {
				brr[j] = -1
				count++
			}

		}

		if brr[i] != -1 {
			brr[i] = count
		}
	}

	for i := 0; i < len(arr); i++ {
		if brr[i] != -1 {
			fmt.Println(arr[i], " ", brr[i])
		}

	}
}

// Hash frequency

func hashFrequency(arr6 []int) {

	brr := make(map[int]int)

	for i := 0; i < len(arr6); i++ {

		_, ok := brr[arr6[i]]

		if !ok {
			brr[arr6[i]] = 1
		} else {
			brr[arr6[i]] = brr[arr6[i]] + 1
		}
	}

	fmt.Println(brr)
}

// Increasing Decreasing

func increasingDecreasing(arr []int) {

	mid := (0 + len(arr) - 1) / 2
	fmt.Println(mid)

	var k int
	var pos int

	for i := 0; i < len(arr); i++ {
		k = arr[i]
		pos = i

		for j := i + 1; j < len(arr); j++ {

			if i <= mid {
				if arr[j] < k {
					k = arr[j]
					pos = j
				}
			} else {
				if arr[j] > k {
					k = arr[j]
					pos = j
				}
			}
		}

		arr[i], arr[pos] = arr[pos], arr[i]
	}

	fmt.Println(arr)
}

// Rotate k items

func rotate(arr8 []int, k int) {

	for k > 0 {

		temp := arr8[0]
		for i := 0; i < len(arr8)-1; i++ {
			arr8[i] = arr8[i+1]

		}
		arr8[len(arr8)-1] = temp
		k--
	}
}

// median

func median(arr9 []int) float64 {
	l1 := arr9[(len(arr9)/2)+1]

	if len(arr9)%2 == 1 {
		return float64(l1)
	} else {
		l2 := arr9[len(arr9)/2]
		return float64((l1 + l2) / 2)
	}
}

// Remove duplicates

func removeDuplicates(arr10 []int, size int) ([]int, int) {

	i := 0
	for j := 1; j < size; j++ {

		if arr10[i] != arr10[j] {
			i++
			arr10[i] = arr10[j]
		}
	}
	return arr10, i + 1
}

// quick sort

func quickSort(arr []int, low int, high int) {

	if low < high {

		pivot := partition(arr, low, high)
		quickSort(arr, pivot+1, high)
		quickSort(arr, low, pivot-1)
	}
}

func partition(arr []int, low int, high int) int {

	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// GCD of 2 numbers

func gcd(num1 int, num2 int) int {
	var small int
	var gcd int
	if num1 < num2 {
		small = num1
	} else {
		small = num2
	}

	for i := 1; i <= small; i++ {
		if num1%i == 0 && num2%i == 0 {
			gcd = i
		}
	}
	return gcd
}

// amrStrong

func armStrong(num int) bool {

	check := 0
	temp := num
	var x int
	for temp > 0 {
		x = temp % 10
		check += x * x * x
		temp = temp / 10
	}

	if check == num {
		return true
	} else {
		return false
	}

}
