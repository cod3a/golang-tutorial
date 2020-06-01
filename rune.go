package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func firstLetterToLowerUsingRunes() {
	var name = "Eisson Alipio"
	var a = []rune(name) //[69 105 115 115 111 110 32 65 108 105 112 105 111]
	//rune = int32
	a[0] = unicode.ToLower(a[0])
	fmt.Println(string(a))
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}
