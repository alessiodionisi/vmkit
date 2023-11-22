package server

import (
	"github.com/alessiodionisi/vmkit/proto"
)

type protoServer struct {
	proto.UnimplementedVMKitServer

	server *Server
}

// func (p *protoServer) Apply(req *proto.ApplyRequest, srv proto.VMKit_ApplyServer) error {
// 	messageChan := make(chan string)
// 	go p.server.engine.Apply(messageChan, req.Data)

// 	for message := range messageChan {
// 		if err := srv.Send(&proto.ApplyResponse{
// 			Message: &message,
// 		}); err != nil {
// 			return status.Error(codes.Internal, err.Error())
// 		}
// 	}

// 	return nil
// }
