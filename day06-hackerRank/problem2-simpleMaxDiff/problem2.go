package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'maxDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY px as parameter.
 */

// maxDifference betweeen 2 value in array
// Calculate the positive difference between the largest and smallest values in the array.
func maxDifference(px []int32) int32 {
	var MaxDiff int32
	min := px[0]

	MaxDiff = -1

	for i := 1; i < len(px); i++ {
		if px[i] > min {
			MaxDiff = int32(math.Max(float64(px[i]-min), float64(MaxDiff)))
		} else {
			min = int32(math.Min(float64(min), float64(px[i])))
		}
	}
	return MaxDiff
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	pxCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var px []int32

	for i := 0; i < int(pxCount); i++ {
		pxItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		pxItem := int32(pxItemTemp)
		px = append(px, pxItem)
	}

	result := maxDifference(px)

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
