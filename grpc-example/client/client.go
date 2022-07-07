package ganeshma_grpc_example

import (
	"context"
	"log"
	"net"
	"time"

	pb "ganeshma.grpc.example/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func RunClient(url string, runforever bool) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to server on port %v", url)
	}
	defer conn.Close()

	c := pb.NewHandlerClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r, err := c.GetName(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Unable to get name from server :: %#+v\n", err.Error())
	}
	log.Print(r.GetName())

	if runforever {
		for {
			r, err := c.GetName(ctx, &emptypb.Empty{})
			if err != nil {
				log.Fatalf("Unable to get name from server :: %#+v\n", err.Error())
			}
			log.Print(r.GetName())
		}
	}
}

func RunClientUnix(socketfile string, runforever bool) {
	dialer := func(addr string, t time.Duration) (net.Conn, error) {
		return net.Dial("unix", socketfile)
	}

	conn, err := grpc.Dial(socketfile, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDialer(dialer)}...)
	if err != nil {
		log.Fatalf("Client failed to connect to socket with error %s\n", err.Error())
	}
	defer conn.Close()

	c := pb.NewHandlerClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r, err := c.GetName(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Unable to get name from server :: %#+v\n", err.Error())
	}
	log.Print(r.GetName())

	if runforever {
		for {
			r, err := c.GetName(ctx, &emptypb.Empty{})
			if err != nil {
				log.Fatalf("Unable to get name from server :: %#+v\n", err.Error())
			}
			log.Print(r.GetName())
		}
	}
}
