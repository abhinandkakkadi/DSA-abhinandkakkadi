package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "abhinand is awesome"
	countVCS(str)

	a := 'a'
	covert := ascii(a)
	fmt.Printf("ascii of %v = %d\n", string(a), covert)

	str2 := "golang is awesome"
	fmt.Println(removeVowels(str2))

	ana1 := "RULES"
	ana2 := "LESRT"
	freq := anagram(ana1, ana2)
	fmt.Println("is anagram : ", freq)
}

func countVCS(str string) {

	vowels := "aeiou"
	countC := 0
	countV := 0
	countS := 0

	for _, val := range str {
		if string(val) == " " {
			countS++
		} else if strings.ContainsAny(string(val), vowels) {
			countV++
		} else {
			countC++
		}
	}

	fmt.Printf("consonants = %d,vowels = %d,spaces = %d\n", countC, countV, countS)
}

func ascii(a rune) int {

	return int(a)

}

func removeVowels(str2 string) string {

	vowels := "aeiou"
	final := ""
	for _, val := range str2 {

		if !strings.ContainsAny(string(val), vowels) {
			final += string(val)
		}
	}

	return final
}

func anagram(ana1 string, ana2 string) bool {

	a1 := make(map[string]int)
	a2 := make(map[string]int)

	for _, val := range ana1 {
		a1[string(val)]++
	}

	for _, val := range ana2 {
		a2[string(val)]++
	}

	flag := 0
	for i, _ := range a1 {

		if a1[i] != a2[i] {
			flag = 1
			break
		}
	}

	if flag == 1 {
		return false
	}

	return true

}
