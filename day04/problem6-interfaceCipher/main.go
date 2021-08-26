package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name       string
	NameEncode string
	Score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *Student) Encode() string {
	var nameEncode = caesarEncrypt(3, s.Name)
	return nameEncode
}

func caesarEncrypt(offset int, input string) string {
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

func asciiToString(ascii []int) string {
	var str string
	for _, v := range ascii {
		str += string(rune(v))
	}
	return str
}

func (s *Student) Decode() string {
	var nameDecode = caesarDecrypt(3, s.NameEncode)
	return nameDecode
}

func caesarDecrypt(offset int, input string) string {
	input = strings.ToLower(input)

	var strAsciiValue []int
	for _, v := range input {
		strAsciiValue = append(strAsciiValue, int(v))
	}

	for i := 0; i < len(strAsciiValue); i++ {
		appendThis := (strAsciiValue[i]-offset-97)%26 + 97
		strAsciiValue[i] = appendThis
	}

	strValue := asciiToString(strAsciiValue)
	return string(strValue)
}

func main() {

	var menu int

	var a = Student{}

	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")

	fmt.Scan(&menu)

	if menu == 1 {

		fmt.Print("\nInput Student’s Name : ")

		fmt.Scan(&a.Name)

		fmt.Print("\nEncode of Student’s Name " + a.Name + " is : " + c.Encode())

	} else if menu == 2 {

		fmt.Print("\nInput Student’s Encode Name : ")

		fmt.Scan(&a.NameEncode)

		fmt.Print("\nDecode of Student’s Name " + a.NameEncode + " is : " + c.Decode())

	} else {

		fmt.Println("Wrong input name menu")

	}

}
