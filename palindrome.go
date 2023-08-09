package main

import (
	"fmt"
	"unicode"
)

func PalindromeTwo(s string) string {
	// Remove non-alphanumeric characters and convert to lowercase
	cleanedStr := ""
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleanedStr += string(unicode.ToLower(r))
		}
	}

	// Compare the cleaned string with its reverse
	length := len(cleanedStr)
	for i := 0; i < length/2; i++ {
		if cleanedStr[i] != cleanedStr[length-1-i] {
			return "false"
		}
	}
	return "true"
}

func main() {
	f := func(i string) {
		fmt.Printf("%v is %v\n", i, PalindromeTwo(i))
	}
	f("Noel - sees Leon")
	f("A war at Tarawa!")
	f("not true!")

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	//fmt.Println(PalindromeTwo(readline()))
}
