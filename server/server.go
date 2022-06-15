package server

import (
	"fmt"
	"net"

	"github.com/adnsio/vmkit/proto"
	"google.golang.org/grpc"
)

type Server struct {
	listener   net.Listener
	grpcServer *grpc.Server
}

func (s *Server) Close() error {
	s.grpcServer.Stop()

	if err := s.listener.Close(); err != nil {
		return fmt.Errorf("net: %w", err)
	}

	return nil
}

func (s *Server) Shutdown() error {
	s.grpcServer.GracefulStop()

	return s.Close()
}

func (s *Server) ListenAndServe() error {
	var err error

	s.listener, err = net.Listen("tcp", "[::1]:8000")
	if err != nil {
		return fmt.Errorf("net: %w", err)
	}

	s.grpcServer = grpc.NewServer()

	protoSrv := &protoServer{}
	proto.RegisterVMKitServer(s.grpcServer, protoSrv)

	if err := s.grpcServer.Serve(s.listener); err != nil {
		return fmt.Errorf("grpc: %w", err)
	}

	return nil
}

func NewServer() *Server {
	return &Server{}
}
