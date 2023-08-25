package main

import (
	"fmt"
	"regexp"
	"strconv"
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

	str3 := "123as"
	fmt.Println(checkStringIsNumber(str3))  


	str4 := "I am a Go developer"
	fmt.Println(removeAllWhiteSpace(str4))

	str5 := `write
					 a multi line
					 string`
	fmt.Println(str5)

	strComp1 := "abhinand"
	strComp2 := "abhinand"
	fmt.Println(Compare(strComp1,strComp2))

	strComp3 := "nandu"
	strComp4 := "abhinandu"
	fmt.Println(Contains(strComp3,strComp4))

	strComp5 := "abhinand Kakkadi Go Developer"
	fmt.Println(split(strComp5))

	str6 := "abhinand kakkadi from kannur"
	fmt.Println(wordsInString(str6))

	str7 := []string{"abhinand","kakkadi","from","haeywa"}
	fmt.Println(joinString(str7))
	
	// Check if string begins with a prefix
	fmt.Println(hasPrefix("abhinand"))

	// Check if string ends with a suffix in Go
	fmt.Println(hasSuffix("abhinand"))


	// to lower case
	fmt.Println(toLowerCase("AbhinandKR"))

	// to upper case 
	fmt.Println(toUpper("abhinandkr"))


	// convert first letter of all words as capital
	fmt.Println(firstLetterOfWordCapital("abhinand kakkadi"))

	// trim prefix
	fmt.Println(trimPrefix("abhinand"))


	// trim suffix
	fmt.Println(trimSuffix("abhinand"))


	//  trim leading and trailing space
	fmt.Println(trimLeadAndTrailSpace(" abhinand kakkadi "))

	// count substring in string
	fmt.Println(countSubStringInString("abhinandabhishekabhinav"))

	// first instance of a substring in string
	fmt.Println(firstInstanceOfSubstring("abhishekabhinav","abhin"))

	// Replace all instance of a substring with another
	fmt.Println(replaceSubstringWithAnother("abhinand","abhi","athi"))

	// Replace first instance of a substring with another
	fmt.Println(replaceFirstSubString("abhinandabhishek","abhi","athi"))

	// find last instance of a substring in a string 
	fmt.Println(indexOfLastInstanceOfSubstring("abhiabhiabhi","abhi"))

	// index string
	indexString("abhinand")


	// reverse string
	reverseString("abhinand")

	// swap two string
	swapToString("abhinand","athira")

	// reverse a string
	fmt.Println(reverseAString("abhinand"))

	// delete a substring
	fmt.Println(deleteSubString("abhinandathira"))

	// delete char at index
	fmt.Println(deleteCharAtIndex("abhinandathira",0))

	// repeat and concatenate a string count times
	fmt.Println(repeatString("abhi",4))

	// case insensitive equal
	fmt.Println(caseInsensitiveEqual("abhi","ABhi"))

	convertToAsci(12,23,34)

	printBackSlash()

	printDoubleQuotes()

	fmt.Println(isMultipleWhiteSpacePresent("abhinand is ok"))

	
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

// prep
func checkStringIsNumber(str3 string) bool {

	_,err := strconv.Atoi(str3)
	return err == nil 
	
}

func removeAllWhiteSpace(str4 string) string {

	trimString := strings.ReplaceAll(str4," ","")
	return trimString

}

// check if 
func Compare(strComp1 string,strComp2 string) bool {

	res := strings.Compare(strComp1,strComp2)
	return  res == 0 
		
}

// check whether the second argument is substring of first argument
func Contains(str1 string,str2 string) bool {

	res := strings.Contains(str2,str1)
	return res

}

func split(strComp5 string) []string {

	// splitting by " "
	// s := strings.Split(strComp5," ")
	// return s

	// splitting by "" - we will get each characters as result
	s := strings.Split(strComp5,"")
	return s


}


func wordsInString(str6 string) []string {

	res := strings.Fields(str6)
	return res
	
}

func joinString(str7 []string) string {

	combinedString := strings.Join(str7,"$")
	return combinedString

}

func hasPrefix(str string) bool {

	res := strings.HasPrefix(str,"abhi")
	return res

}

func hasSuffix(str string) bool {

	res := strings.HasSuffix(str,"nand")
	return res

}

func toLowerCase(str string) string {

	res := strings.ToLower(str)
	return res

}

func toUpper(str string) string {

	return strings.ToUpper(str)

}

func firstLetterOfWordCapital(str string) string {

	return strings.Title(str)

}

func trimPrefix(str string) string {
	return strings.TrimPrefix(str,"abhi")
}

func trimSuffix(str string) string {

	return strings.TrimSuffix(str,"nand")

}

func trimLeadAndTrailSpace(str string) string {

	return strings.TrimSpace(str)

}

func countSubStringInString(str string) int {

	return strings.Count(str,"abhi")

}

func firstInstanceOfSubstring(str string,sub string) int {

	return strings.Index(str,sub)

} 

func replaceSubstringWithAnother(str string,old string,new string) string {

	return strings.ReplaceAll(str,old,new)

}

func replaceFirstSubString(str string,old string,new string) string {

	return strings.Replace(str,old,new,2)
}

func indexOfLastInstanceOfSubstring(str string,sub string) int {

	return strings.LastIndex(str,sub)

}

func indexString(str string) {

	a := []rune(str)

	// when ranging through string - it's by default a string
	// we have to use rune while character printing as string is stored as a sequence of bytes
	// and a single character can be represented by more than 1 byte if it is a foreign character other than english.
	for _,val := range a {
		fmt.Printf("%c",val)
	}

}

func reverseString(str string) string {

	str2 := []rune(str)

	for i,j := 0,len(str2)-1; i <=j; i,j = i+1,j-1 {
		str2[i],str2[j] = str2[j],str2[i]
	}

	return string(str2)

}

func swapToString(str1 string,str2 string) {

	fmt.Println("string 1 :",str1)
	fmt.Println("string 2 :",str2)

	str1,str2 = str2,str1
	fmt.Println("string 1 :",str1)
	fmt.Println("string 2 :",str2)
}	


func reverseAString(str string) string {

	str2 := []rune(str)

	str3 := []rune{}

	for i := len(str2) - 1; i >= 0; i-- {

		str3 = append(str3,str2[i])

	}

	return string(str3)
}

func deleteSubString(str string) string {

	// here we are deleting that parti
	// same thing can be done to delete a single character too
	return strings.ReplaceAll(str,"a","")

}

func deleteCharAtIndex(str string,index int) string {

	str2 := []rune(str)

	str2 = append(str2[:index],str2[index+1:]... )
	return string(str2)

}

func repeatString(str string,count int) string {

	return strings.Repeat(str,count)

}

func caseInsensitiveEqual(str1 string,str2 string) bool {

	return strings.EqualFold(str1,str2)

}

func convertToAsci(nums ...int) {

	for _, digit := range nums {
		fmt.Printf("%s",string(digit))
	}
}


func printBackSlash() {

	fmt.Println("\\abhinand")

	fmt.Println(`\abhinand`)
}


func printDoubleQuotes() {

	fmt.Println("\"abhinand\"")

	fmt.Println(`"abhinand"`)

}


func isMultipleWhiteSpacePresent(str string) bool {

	present := regexp.MustCompile(`\s`).MatchString(str)
	return present

}