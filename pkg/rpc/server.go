// Virtual Machine manager that supports QEMU and Apple virtualization framework on macOS
// Copyright (C) 2021 VMKit Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
