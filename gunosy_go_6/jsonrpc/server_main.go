package main

import (
	"github.com/hkurokawa/go-sandbox/gunosy_go_6/data"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	arith := new(data.Arith)

	rpc.Register(arith)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go jsonrpc.ServeConn(conn)
	}
}
