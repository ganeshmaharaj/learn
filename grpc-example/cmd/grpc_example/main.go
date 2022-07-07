package main

import (
	"flag"
	"fmt"

	client "ganeshma.grpc.example/client"
	server "ganeshma.grpc.example/server"
)

const (
	defaultPort   = 50051
	defaultSocket = "/tmp/ganeshma_grpc_example.socket"
)

func main() {

	flagServer := flag.Bool("server", false, "Start the server")
	flagClient := flag.Bool("client", false, "run the client part of the code")
	flagRunLong := flag.Bool("forever", false, "Run client for long")
	flagHost := flag.String("host", "localhost", "Host to connect to")
	flagUnixSocket := flag.Bool("unix", false, "Use unix socket file for communication")
	//flagParallel := flag.Uint64("parallel", 1, "Default number of parallel calls to make")

	flag.Parse()

	if *flagServer {
		if *flagUnixSocket {
			server.RunServerUnix(defaultSocket)
		} else {
			server.RunServer(defaultPort)
		}
	}

	if *flagClient {
		if *flagUnixSocket {
			client.RunClientUnix(defaultSocket, *flagRunLong)
		} else {
			client.RunClient(fmt.Sprintf("%s:%d", *flagHost, defaultPort), *flagRunLong)
		}
	}
}
