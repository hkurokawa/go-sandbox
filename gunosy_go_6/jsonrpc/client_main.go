package main

import (
	"fmt"
	"github.com/hkurokawa/go-sandbox/gunosy_go_6/data"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer conn.Close()
	client := jsonrpc.NewClient(conn)

	// Synchronous call
	args := &data.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}
