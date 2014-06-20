package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const input = "6153 746c 6465 5f5f f9b2 0b18 31ed 9cd9 9b88 b6d9 4bfc 8cba 4a06 5118 e8af c310 2303 c17e 349b 08a8 61a3 025f db53 e2c4"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	var buff bytes.Buffer
	for scanner.Scan() {
		tok := scanner.Text()
		upper := tok[:2]
		lower := tok[2:]
		buff.Write(parseHex(lower))
		buff.Write(parseHex(upper))
	}
	buff.WriteTo(os.Stdout)
}

func parseHex(input string) []byte {
	v, err := strconv.ParseInt(input, 16, 16)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse the string %s as 16-base integer: %s", input, err)
		return []byte{0}
	}
	return []byte{byte(v)}
}
