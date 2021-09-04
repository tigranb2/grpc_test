package roles

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/msg"
	"time"
)

var acks = 3
var commits float64 = 0
var responses = make(chan *msg.Msg)

type connection struct {
	client msg.MessengerClient
	m      chan *msg.Msg
}

func Sender(ips []string) {
	connections := []connection{}
	for _, ip := range ips {
		if c := dialServer(ip); c != nil {
			connections = append(connections, connection{client: c, m: make(chan *msg.Msg)})
		}
	}

	for i := 0; i < len(connections); i++ { //start unicast
		go connections[i].unicast()
	}

	start := time.Now()
	for i := 0; i < len(connections); i++ { //gives unicast message to send
		connections[i].m <- &msg.Msg{Id: int32(commits)}
	}

	for commits < 10000 {
		if acks >= 2 {
			commits++
			acks = 0

			for i := 0; i < len(connections); i++ {
				connections[i].m <- &msg.Msg{Id: int32(commits)}
			}
		}

		select {
		case resp := <-responses:
			if resp.Id == int32(commits) {
				acks++
			}
		}
	}

	duration := time.Since(start)
	fmt.Println(commits / duration.Seconds())
}

func dialServer(ip string) msg.MessengerClient {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*10)
	conn, err := grpc.DialContext(ctx, ip, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		return nil
	}

	return msg.NewMessengerClient(conn)
}

func (c *connection) unicast() {
	stopUnicast := make(chan bool)
	stream, err := c.client.SendMsg(context.Background())
	if err != nil {
		stopUnicast <- true
	}

	//receive messages
	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				stopUnicast <- true
			}

			responses <- resp
		}
	}()

	//send messages
	go func() {
		for {
			select {
			case req := <-c.m:
				if err := stream.Send(req); err != nil {
					stopUnicast <- true
				}
			}
		}
	}()

	<-stopUnicast
}
