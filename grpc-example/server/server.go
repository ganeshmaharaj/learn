package ganeshma_grpc_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "ganeshma.grpc.example/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedHandlerServer
}

func (s *server) GetName(ctx context.Context, _ *emptypb.Empty) (*pb.NameResponse, error) {
	var hName string
	var err error
	if hName, err = os.Hostname(); err != nil {
		log.Fatal("Unable to get hostname")
	}
	log.Print("Received request")
	log.Printf("%#+v\n", ctx)
	return &pb.NameResponse{Name: hName}, nil
}

func RunServer(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Printf("Unable to listen to port %v", port)
	}
	s := grpc.NewServer()
	pb.RegisterHandlerServer(s, &server{})
	log.Printf("Server listening to port %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server at port %v", port)
	}
}
