package rpc

import (
	"io"
	"net/rpc"

	"github.com/adnsio/vmkit/pkg/engine"
)

type NewServerOptions struct {
	Engine *engine.Engine
}

type Server struct {
	server *rpc.Server
}

func (s *Server) ServeConn(conn io.ReadWriteCloser) {
	s.server.ServeConn(conn)
}

func NewServer(opts *NewServerOptions) (*Server, error) {
	server := rpc.NewServer()

	server.RegisterName("VirtualMachine", &VirtualMachineReceiver{
		engine: opts.Engine,
	})

	return &Server{
		server: server,
	}, nil
}
