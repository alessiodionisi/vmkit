package main

import (
	"fmt"

	"github.com/adnsio/vmkit/proto"
	"google.golang.org/grpc"
)

func newClient() (*grpc.ClientConn, proto.VMKitClient, error) {
	conn, err := grpc.Dial("[::1]:8000", grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("client error: %w", err)
	}

	client := proto.NewVMKitClient(conn)
	return conn, client, nil
}
