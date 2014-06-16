package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// open input file
	in, err := os.Open("jp.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := in.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	r := bufio.NewReader(in)

	// http://codepoints.net/U+672C "本"
	rune, size, _ := r.ReadRune()
	fmt.Printf("The first rune: %x, size=%d\n", rune, size)

	if err := r.UnreadRune(); err != nil {
		panic(err)
	}
	fmt.Printf("The first rune again: %x, size=%d\n", rune, size)
}
