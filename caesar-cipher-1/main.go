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
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
	// Write your code here
	k = k % 26

	var ret string
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			tmp := (c + k) % ('z' + 1)
			if tmp < 'a' {
				tmp += 'a'
			}
			ret += string(tmp)
		} else if c >= 'A' && c <= 'Z' {
			tmp := (c + k) % ('Z' + 1)
			if tmp < 'A' {
				tmp += 'A'
			}
			ret += string(tmp)
		} else {
			ret += string(c)
		}
	}
	return ret
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	_ = int32(nTemp)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

	fmt.Fprintf(writer, "%s\n", result)

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
