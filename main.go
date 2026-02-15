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
