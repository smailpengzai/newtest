package main

import (
	"github.com/smallnest/rpcx/server"
)

func main() {
	s := server.NewServer()
	s.RegisterName("Arith", new(models.Arith), "")
	//s.Register(new(example.Arith), "")
	s.Serve("tcp", ":8972")
}
