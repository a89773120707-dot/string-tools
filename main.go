package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("== string-tools ==")
	fmt.Println("hello, go 'git- traning! ->", CleanString("hello, go 'git- traning! "))
	fmt.Println("учение       свет   !", " -> ", RemoveSpaces("учение       свет   !"))

	fmt.Println("'radar' is palindrome -> ", IsPalindrome("radar"))
}

func CleanString(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))

	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func RemoveSpaces(s string) string {

	words := strings.Fields(s)

	if len(words) == 0 {
		return ""
	}
	return strings.Join(words, " ")

}

func IsPalindrome(s string) bool {
	runes := []rune(s)

	left, right := 0, len(runes)-1

	for left < right {
		if !unicode.IsLetter(runes[left]) || !unicode.IsDigit(runes[left]) {
			left++
		}
		if !unicode.IsLetter(runes[right]) || !unicode.IsDigit(runes[right]) {
			right--
		}
		if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
			return false
		}
		left++
		right--
	}
	return true
}
