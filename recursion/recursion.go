package main

import (
	"fmt"
	"math"
	"unicode"
)

func main() {

	// fibanochi of a number
	fmt.Println(fib(5))

	// implement binary search  using recursion
	arr := []int{1,2,3,4,5,6,7,8,9}
	search := bSearch(arr,0,len(arr)-1,4)
	if search != -1 {
		fmt.Println("Element is present at the position")
	}

	// max and min element of an array using recursion
	min := math.MaxInt64
  max := math.MinInt64
	
	min,max := maxMin(arr,0,len(arr)-1,max,min)
	fmt.Println("Max and Min value inside the array are max:= ",max," min:= ",min)

	// Find first capital letter in an array
	str
}


// fibanochi series up to a given number
func fib(n int) int {

	if n==0 || n==1 {
		return n;
	}

	return fib(n-1) + fib(n-2)
 
 }


// implement binary  search using recursion
 func bSearch(arr []int,target int,s int,e int,max int,min int) int {

	if s > e {
		return -1
	}
	mid := (s+e)/2
 
 
	if target == arr[mid] {
		return mid;
	}
 
 
	if target > arr[mid] {
		s = mid +1
	} else {
		e = mid -1
	}

	return bSearch(arr,target,s,e)

	}
 

	// find minimum and maximum value in an array using recursion
	func maxMin(arr []int,s int, e int, min int, max int) (int,int) {

		if(s == e+1) {
			return min,max
		}
	 
		if arr[s] > max {
			max = arr[s]
		}
	 
		if arr[s] < min {
			min = arr[s]
		}
	 
		return maxMin(arr,s+1,e,min,max)
		
	 }
	 
 
	//  find first capital letter in a  string
	func capital(str string, s int, l int) string {
		if( s==l+1 ) {
			return ""
		}
	 
	 
		if( unicode.IsUpper(rune(str[s]))) {
			return string(str[s])
		}
	 
	 
		return capital(str,s+1,l)
	 }
	 
	//  Reverse a string using recursion

	func reverse(str []byte, s int,l int) []byte {

		if s > l {
			return str
		}
	 
		temp := str[s]
		str[s] = str[l];
		str[l] = temp;

		return reverse(str,s+1,l-1)
	 }
	 

	//  Print numbers from 1 - N using recursion
	func reverse(n int,size int) {


		if n == size {
			fmt.Println(n)
			return
		}
	 
	 
		fmt.Println(n);
		reverse(n+1,size)
	 
	 
	 }
	 