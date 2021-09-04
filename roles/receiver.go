package roles

import (
	"google.golang.org/grpc"
	"grpc_test/msg"
	"log"
	"net"
)

type server struct {
	msg.UnimplementedMessengerServer
}

func Receiver(connections []string, id int) {
	ip := connections[id-1]
	s := server{}
	ln, err := net.Listen("tcp", ip)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	grpcServer := grpc.NewServer()
	msg.RegisterMessengerServer(grpcServer, &s)
	if err := grpcServer.Serve(ln); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}

func (s *server) SendMsg(stream msg.Messenger_SendMsgServer) error {
	errStatus := make(chan error)

	go func() {
		for {
			rec, err := stream.Recv() //reads sender messages
			if err != nil {
				errStatus <- err
			}

			if err := stream.Send(rec); err != nil { //sends message back to sender as ACK
				errStatus <- err
			}

		}
	}()

	return <-errStatus
}
