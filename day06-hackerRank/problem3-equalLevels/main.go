package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'updateTimes' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY signalOne
 *  2. INTEGER_ARRAY signalTwo
 */

// Check how many element of the two array are equal
func updateTimes(signalOne []int32, signalTwo []int32) int32 {
	// Write your code here
	var count int32
	var currMaxValue int32

	currMaxValue = -1

	// Cut the longest array to the same length
	if len(signalOne) > len(signalTwo) {
		signalOne = signalOne[:len(signalTwo)]
	} else if len(signalTwo) > len(signalOne) {
		signalTwo = signalTwo[:len(signalOne)]
	}

	for i := 0; i < len(signalOne); i++ {
		if signalOne[i] == signalTwo[i] {
			if signalOne[i] <= currMaxValue {
				continue
			} else {
				currMaxValue = signalOne[i]
				count++
			}
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	signalOneCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var signalOne []int32

	for i := 0; i < int(signalOneCount); i++ {
		signalOneItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		signalOneItem := int32(signalOneItemTemp)
		signalOne = append(signalOne, signalOneItem)
	}

	signalTwoCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var signalTwo []int32

	for i := 0; i < int(signalTwoCount); i++ {
		signalTwoItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		signalTwoItem := int32(signalTwoItemTemp)
		signalTwo = append(signalTwo, signalTwoItem)
	}

	result := updateTimes(signalOne, signalTwo)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
