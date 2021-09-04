package main

import (
	"fmt"
	"grpc_test/roles"
	"log"
	"os"
	"strconv"
)

var connections = []string{"127.0.0.1:8000", "127.0.0.1:8001", "127.0.0.1:8002"}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Please specify node type (r, s) and node id")
		return
	}

	switch arguments[1] {
	case "r":
		nodeId, err := strconv.Atoi(arguments[2])
		if err != nil {
			log.Fatalf("node id formatted incorrectly: %v", err)
		}
		roles.Receiver(connections, nodeId)
	case "s":
		roles.Sender(connections)
	}

}
