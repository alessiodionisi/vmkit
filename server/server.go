package server

import (
	"log/slog"
	"net"

	"github.com/alessiodionisi/vmkit/engine"
	"github.com/alessiodionisi/vmkit/proto"
	"google.golang.org/grpc"
)

type Server struct {
	engine      *engine.Engine
	grpcServer  *grpc.Server
	netListener net.Listener
	logger      *slog.Logger
}

func (s *Server) ListenAndServe(network string, address string) error {
	var err error
	s.netListener, err = net.Listen(network, address)
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer()
	proto.RegisterVMKitServer(s.grpcServer, &protoServer{server: s})

	return s.grpcServer.Serve(s.netListener)
}

func (s *Server) Shutdown() error {
	s.logger.Debug("shutting down server")

	s.grpcServer.GracefulStop()
	return s.netListener.Close()
}

func New(engine *engine.Engine, logger *slog.Logger) *Server {
	return &Server{
		engine: engine,
		logger: logger,
	}
}
