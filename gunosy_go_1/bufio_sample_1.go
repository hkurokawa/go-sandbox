package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	// open input file
	in, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	// close in on exit and check for its returned error
	defer func() {
		if err := in.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	r := bufio.NewReader(in)

	// open output file
	out, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	// close out on exit and check for its returned error
	defer func() {
		if err := out.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(out)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err = w.Flush(); err != nil {
		panic(err)
	}
}
