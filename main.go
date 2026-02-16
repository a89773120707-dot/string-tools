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

	fmt.Println(NormalizePhone("+7 (999) 123-45-67")) // +79991234567
	fmt.Println(NormalizePhone("8 999 123 45 67"))    // +79991234567
	fmt.Println(NormalizePhone("9991234567"))         // +79991234567
	fmt.Println(NormalizePhone("abc123"))             // abc123 (не изменили)

	fmt.Println(TruncateWithEllipsis("Hello world this is test", 10)) // Hello wor...
	fmt.Println(TruncateWithEllipsis("short", 10))                    // short
	fmt.Println(TruncateWithEllipsis("very long string", 3))          // ...
	fmt.Println(TruncateWithEllipsis("", 5))                          // ""

	fmt.Println("alex@example.com ->", maskemail("alex@example.com"))
	fmt.Println("a@example.com ->", maskemail("a@example.com"))
	fmt.Println("al@example.com ->", maskemail("al@example.com"))
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

func NormalizePhone(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	digits := b.String()
	n := len(digits)

	switch {
	case n == 11 && digits[0] == '8':
		return "+7" + digits[1:]
	case n == 11 && digits[0] == '7':
		return "+" + digits
	case n == 10:
		return "+7" + digits
	default:
		return s
	}

}

func TruncateWithEllipsis(s string, maxlen int) string {
	runes := []rune(s)

	if len(runes) <= maxlen {
		return string(runes)
	}

	if maxlen <= 3 {
		return "..."
	}

	return string(runes[:maxlen-3]) + "..."
}

func maskemail(s string) string {
	atIndex := strings.Index(s, "@")

	username := s[:atIndex]
	domain := s[atIndex:]

	runes := []rune(username)
	n := len(runes)

	switch {
	case n == 1:
		return "*" + domain
	case n == 2:
		return string(runes[0]) + "*" + domain
	default:
		stars := strings.Repeat("*", n-2)
		return string(runes[0]) + stars + string(runes[n-1]) + domain
	}

}
