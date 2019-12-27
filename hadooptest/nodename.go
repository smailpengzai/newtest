package main

import (
	"context"
	hadoop_hdfs "github.com/alberts/gohadoop/hdfs"
	albhdfsnode "github.com/alberts/gohadoop/hdfs/namenode"
	"github.com/smallnest/rpcx/client"
	"log"
)

func main() {
	// #1
	d := client.NewZookeeperDiscovery("tcp@", "", "", nil)
	// #2
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// #3
	args := &albhdfsnode.GetBlocksResponseProto{
		Blocks: nil,
	}

	// #4
	reply := &hadoop_hdfs.BlocksWithLocationsProto{}

	// #5
	err := xclient.Call(context.Background(), "GetBlocks", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
