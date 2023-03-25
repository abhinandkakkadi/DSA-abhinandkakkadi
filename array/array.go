package main

import (
	"fmt"
)
func main() {


 arr := make([]int,20)
 fmt.Println("Enter values into the array")
 size := 5

//  Add values into the array
 addArray(arr,&size)
 displayArray(arr,&size)

//  insert element in a sorted array in correct position
 insertArray(arr,&size,4,8)
 displayArray(arr,&size)

// append values into an array
 append(arr,&size,10)
 displayArray(arr,&size)

//  delete an element in a specified index
 deleteElement(arr,&size,9)
 displayArray(arr,&size)

//  reverse an array
 reverseArray(arr,&size)
 displayArray(arr,&size)

// Left rotating
 leftRotating(arr,&size)
 displayArray(arr,&size)

// Left shifting
 leftShifting(arr,&size)
 displayArray(arr,&size)

// right shifting
 rightShifting(arr,&size)
 displayArray(arr,&size)

// right rotation
 rightRotation(arr,&size)
 displayArray(arr,&size)

// Insert in sorted order
 insertInSorted(arr, &size, 3 )
 displayArray(arr,&size)

// check whether an array is sorted
 arraySorted(arr,&size)

// arrange -ve numbers tp left and +ve numbers to right
 arrangeLeft(arr,&size)
 displayArray(arr,&size)


}

// Add elements into an array
func addArray(arr []int,size *int) {
 for i:=0; i<*size; i++ {
   fmt.Scanf("%d",&arr[i]);
 }
}

// display array
func displayArray(arr []int,size *int) {
 fmt.Println("The values of the array")
 for i:=0; i<*size; i++ {
   fmt.Printf("%d\n",arr[i]);
 }
}

// insert element at a specified index
func insertArray(arr []int,size *int, index int,value int) {
 for j := *size; j > index; j-- {
   arr[j] = arr[j-1]
 }
 arr[index] = value
 *size++
}

// append values into the array
func append(arr []int, size *int, value int) {
   arr[*size] = value
   *size++


}

// delete element of the index specified
func deleteElement(arr []int, size *int, index int) {
  for i:=0;i<*size; i++ {
   if i == index {
     for j:=i; j< *size-1; j++ {
       arr[j] = arr[j+1]
     }
     *size = *size -1
   }
 }
}


// Reverse Array
func reverseArray(arr []int, size *int) {
 var temp int
 for i,j := 0,*size-1; i<j; i,j = i+1,j-1 {
   temp = arr[i]
   arr[i] = arr[j]
   arr[j] = temp
 }
}


// left rotating
func leftRotating(arr []int, size *int) {
 first := arr[0]


 for i:=0; i<*size-1;i++ {
   arr[i] = arr[i+1]
 }
 arr[*size-1] = first
}


// left shifting
func leftShifting(arr []int, size *int) {
 for i:=0; i<*size-1; i++ {
   arr[i] = arr[i+1]
 }
 arr[*size-1] = 0
}


// right shifting
func rightShifting(arr []int, size *int) {
 for i:=*size-1; i>0; i-- {
   arr[i] = arr[i-1]
 }
 arr[0] = 0
}

// right rotation
func rightRotation(arr []int, size *int) {
 start := arr[*size-1]


 for i:=*size-1; i>0; i-- {
   arr[i] = arr[i-1]
 }


 arr[0] = start
}


// insert in order
func insertInSorted(arr []int, size *int, value int) {
 var i int
 for i= *size; arr[i-1] > value; i-- {
   arr[i] = arr[i-1]
 }
 arr[i] = value
 *size++
}


// check whether array is sorted
func arraySorted(arr []int,size *int) {
 flag :=0
 for i:=0; i<*size-1; i++ {
   if arr[i] > arr[i+1] {
     flag =1
   }
 }
 if flag == 0 {
   fmt.Println("array sorted")
 } else {
   fmt.Println("Array not sorted")
 }
}


// arranging -ve on left side
func arrangeLeft(arr []int,size *int) {
 index :=0;
 var temp int
 for i:=1; i<*size;i++ {
   if arr[i] < 0 {
     temp = arr[i]
     arr[i] = arr[index]
     arr[index] = temp
     index++
   }
 }
}
