package grpc

import (
	"google.golang.org/grpc"
	"net"
)

func UseGrpc(address string) error {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	server := grpc.NewServer()

	defer func() {
		server.Stop()
		listen.Close()
	}()

	err = server.Serve(listen)
	if err != nil {
		return err
	}
	return nil
}

func StartGrpcServer() {
	go UseGrpc("0.0.0.0:8081")
}
