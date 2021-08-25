package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(caesar3(3, "abc"))
	fmt.Println(caesar3(1000, "abcdefghijklmnopqrstuvwxyz"))
}

func caesar3(offset int, input string) string {
	input = strings.ToLower(input)

	var strAsciiValue []int
	for _, v := range input {
		strAsciiValue = append(strAsciiValue, int(v))
	}

	for i := 0; i < len(strAsciiValue); i++ {
		appendThis := (strAsciiValue[i]+offset-97)%26 + 97
		strAsciiValue[i] = appendThis
	}

	strValue := asciiToString(strAsciiValue)
	return string(strValue)
}

// function from ascii to string
func asciiToString(ascii []int) string {
	var str string
	for _, v := range ascii {
		str += string(rune(v))
	}
	return str
}

func encrypt(r rune, shift int) rune {
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}

func caesar(offset int, input string) string {
	result := strings.Map(func(r rune) rune {
		return encrypt(r, offset)
	}, input)

	return result
}

func ceaser2(offset int, input string) {
	for i := 0; i < len(input); i++ {
		fmt.Printf("%c", input[i]+byte(offset))
	}
}
