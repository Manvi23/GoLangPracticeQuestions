//Find the smallest window in a string containing all characters of another string
//it will return first matching shortest substring
//even if there are more substrings of same size
//it can be easily altered , saved values can be updated as per new values
package main

import (
	"fmt"
	"strings"
)

func CheckContains(s []rune, str string) bool {
	for _, v := range s {
		ok := strings.Contains(str, string(v))
		if !ok {
			return false
		}
	}
	return true
}

func FindSmallestSubstring(s []rune, str string) {
	lenstr := len(str)
	lenrune := len(s)
	if lenstr < lenrune {
		fmt.Println("Sorry ! String is too small to find all chars")
		return
	}
	i := 0
	savedlen := 0
	savedstr := ""
	for i <= lenstr-2 {
		j := i + 2
		for j < lenstr {
			currentSlice := str[i : j+1]
			val := CheckContains(s, string(currentSlice))
			if val {
				currentstr := string(str[i : j+1])
				currentlen := len(str[i:j])
				if savedlen == 0 {
					savedlen = currentlen
					savedstr = currentstr
				}
				if currentlen < savedlen {
					savedlen = currentlen
					savedstr = currentstr
				} else {
					break
				}
			}
			j++
		}
		i++
	}
	if savedstr != "" {
		fmt.Println("Found the Smallest Substring------->", savedstr)
	} else {
		fmt.Println("Nope ! Could not Find")
	}
}

func main() {
	fmt.Println("In main function")
	//s := []rune{'x', 'y', 'z'}
	//str := "xyyyzyzyyx"
	//str := "xxxxyyyyxxxzayyxazzzz"
	//str := "aaaaa"
	//str := "rt"
	//str := "kdfkdsjfkjfdsk"
	//str := "xyyz"
	//str := "xyz"
	//str := "xy"
	//str := "abcdefghizyyz"
	//str := "abcdefghizzyxy"
	//str := "xxyz"

	//another set of inputs
	s := []rune{'a', 'b', 'c'}
	str := "aaabbbcccccaab"

	FindSmallestSubstring(s, str)
}
