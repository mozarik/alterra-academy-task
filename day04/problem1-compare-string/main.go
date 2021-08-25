package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))

	// Test fail cases
	fmt.Println(Compare("KANGOORO", "ZEIN"))
	fmt.Println(Compare("Kang", "KANG"))
}

func Compare(a, b string) string {
	// check if 'a' is a substring of 'b'
	if strings.Contains(b, a) {
		return a
	} else if strings.Contains(a, b) {
		return b
	} else {
		return "Ehe te nandayo"
	}
}
