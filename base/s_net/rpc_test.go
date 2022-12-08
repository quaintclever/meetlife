package snet

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"testing"
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func TestRpcServer(t *testing.T) {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("err :%v", err)
	}
	http.Serve(l, nil)
}

func TestRpcClient(t *testing.T) {
	c, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("err: %v", err.Error())
	}

	args := &Args{
		A: 2,
		B: 3,
	}
	var rs int

	err2 := c.Call("Arith.Multiply", args, &rs)
	if err2 != nil {
		log.Fatalf("err: %v", err2.Error())
	}
	fmt.Printf("Arith: %d * %d = %d \n", args.A, args.B, rs)
}
