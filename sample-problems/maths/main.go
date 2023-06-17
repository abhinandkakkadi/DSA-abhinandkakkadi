package main

import "fmt"

func main() {

	// count digits in a number
	// APPROACH - we will remove each digit from the give number and take count of the removed digit
	// In our case, it's easy to delete element from last, so we wil do that.
	num := 43422
	count := sumOfDigits(num)
	fmt.Println("the count of number is : ", count)

	// Check whether a number is palindrome or not
	// APPROACH : I will take in the last digit and make it the first digit and second lsat digit as second digit and so on
	// at the end if the manipulated number and original number is the same, we can say that it is a palindrome number
	num2 := 12221
	isPal := isPalindrome(num2)
	if isPal {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}

	// factorial of a number
	// APPROACH : if the number i 0 or 1 then 1 is the factorial
	// other wise iterate from 2 to n and multiply all the numbers along the way to a variable which was initially initialized to 1
	num3 := 5
	r := fact(num3)
	fmt.Println("factorial of a number", r)
}

// sum of digits
func sumOfDigits(num int) int {

	count := 0
	for num > 0 {
		count++
		num = num / 10
	}
	return count

}

// check whether the given number is palindrome
func isPalindrome(num int) bool {

	rev := 0
	temp := num

	for temp > 0 {
		rev = rev*10 + temp%10
		temp = temp / 10
	}

	if rev == num {
		return true
	}

	return false
}

// factorial of a number
func fact(num int) int {

	if num == 0 || num == 1 {
		return num
	}

	res := 1
	for i := 2; i <= num; i++ {
		res = res * i
	}

	return res
}
