package main

import "fmt"

func main() {

	c := countSetBits(10)
	fmt.Println(c)

	c2 := countSetBrain(10)
	fmt.Println(c2)

	// power of two

	fmt.Println(powerOfTwo(16))
}

// method 1
func countSetBits(n int) int {

	res := 0

	for n > 0 {
		res = res + (n & 1)
		n = n / 2
	}

	return res

}

// method 2
func countSetBrain(n int) int {

	res := 0

	for n > 0 {
		n = n & (n - 1)
		res++
	}

	return res
}

//  power of two

func powerOfTwo(n int) bool {

	if n <= 0 {
		return false
	}

	// count := 0

	// for n > 0 {
	// 	n = n & (n-1)
	// 	count++
	// }

	// if count == 1 {
	// 	return true
	// }
	// return false

	return (n & (n - 1)) == 0
}
