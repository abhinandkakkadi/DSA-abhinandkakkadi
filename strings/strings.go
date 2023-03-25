package main


import (
 "fmt"
 "strings"
 "unicode"
)


func main() {
 word := "malayalam is the best"

 // length of the string
 length(word)

 // convert string to upper case
 uppercase(word)

 // check whether a string is palindrome
 palindrome(word)

 // remove vowels from the strings
 removeVowels(word)

 // number of words in a string
 numberWords(word)

 // Reverse words in an array
 reverseWords(word)

}


// length of string
func length(word string) {
 len :=0
 arr := []byte(word)
 for range arr {
   len++
 }
 fmt.Println("Length of the string is =",len)
}


// convert to uppercase
func uppercase(word string) {
 var ch []rune
 for _,val := range  word {
   ch = append(ch, unicode.ToUpper(val))
 }
 fmt.Println(string(ch))
}


// check whether the string is palindrome
func palindrome(word string) {
 flag :=0
 for i,j:=0,len(word)-1; i<j; i,j=i+1,j-1 {
   if word[i] != word [j] {
     flag =1
   }
 }
 if flag == 0 {
   fmt.Println("palindrome")
 } else {
   fmt.Println("Not paliandrome")
 }
}


// Remove Vowels from a string
func removeVowels(word string) {
 word2 := ""
 for _,val := range word {
   if string(val) != "a" && string(val) != "e" &&  string(val) != "i" && string(val) != "o" && string(val) != "u" {
     word2+=string(val)
   }
 }
 fmt.Println(word2)
}


// Number of words in a string
func numberWords(word string) {
 count :=0
 w := false
 for _,val := range word {
   if val == ' ' {
     if w {
       count++
       w = false
     }
   } else {
     w = true
   }
 }
 fmt.Println("Number of words: ",count+1)
}


// Reverse words of a string
func reverseWords(word string) {


 words := strings.Fields(word)
  for i,j := 0,len(words)-1; i<j; i,j = i+1,j-1 {
   words[i],words[j] = words[j],words[i]
 }
 str := strings.Join(words," ")
 fmt.Println(str)
}


