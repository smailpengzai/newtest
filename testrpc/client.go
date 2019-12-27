package main

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
)

type Argsc struct {
	A int
	B int
}

type Replyc struct {
	C int
}

type Arithc int

func (t *Arithc) Mulc(ctx context.Context, args *Argsc, reply *Replyc) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	// #1
	d := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:8972", "")
	// #2
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// #3
	args := &Argsc{
		A: 10,
		B: 20,
	}

	// #4
	reply := &Replyc{}

	// #5
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
