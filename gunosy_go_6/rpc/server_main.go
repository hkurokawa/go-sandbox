package main

import (
	"github.com/hkurokawa/go-sandbox/gunosy_go_6/data"
	"log"
	"net/http"
	"net/rpc"
)

func main() {
	arith := new(data.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal(err)
	}
}
