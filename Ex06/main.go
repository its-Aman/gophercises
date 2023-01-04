package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// var input string
	// fmt.Scanf("%s\n", &input)
	// words := camelProblem(input)
	// fmt.Println("# of Words: ", words)

	originalText := "middle-Outz"
	cipheredText := caesarCipher(originalText, 2)
	fmt.Printf("Original is: %s\nCiphered Text is: %s\n", originalText, cipheredText)
}

func camelProblem(input string) int {
	ans := 1

	for _, ch := range input {
		s := string(ch)
		if s == strings.ToUpper(s) {
			ans++
		}
	}

	return ans
}

func caesarCipher(s string, k int32) string {
	ans := ""

	for _, c := range s {

		if !unicode.IsLetter(c) {
			ans += string(c)
			continue
		}

		base := 'a'

		if unicode.IsUpper(c) {
			base = 'A'
		}

		ans += string((((c - base) + k) % 26) + base)
	}

	return ans
}
