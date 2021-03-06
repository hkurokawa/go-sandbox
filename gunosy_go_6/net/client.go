package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Connet to TCP port 2000 on all interfaces.
	l, err := net.Dial("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	fmt.Fprintf(l, "Hello, world!\n")
	response, err := bufio.NewReader(l).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
