package serverDomain

import (
	"context"
	"fmt"
	"gfs-go/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ChunkServer struct {
	ServerAddr        string
	ChunkIds          []string
	MemoryUtilization float32
}

type Server struct {
	pb.MasterServiceServer
	ChunkServers []ChunkServer
}

var MasterServer Server
var ReplicationFactor = 3

func (s *Server) HeartBeat(ctx context.Context, req *pb.HeartBeatRequest) (*emptypb.Empty, error) {
	chunkServer := &ChunkServer{
		ServerAddr:        req.Addr,
		ChunkIds:          req.ChunkIds,
		MemoryUtilization: req.MemoryUtilization,
	}
	s.ChunkServers = addIfNotPresent(s.ChunkServers, *chunkServer)
	MasterServer = *s
	fmt.Println(req.Addr)
	return &emptypb.Empty{}, nil
}

func addIfNotPresent(c []ChunkServer, element ChunkServer) []ChunkServer {
	for i, v := range c {
		if v.ServerAddr == element.ServerAddr {
			c[i] = element
			return c
		}
	}
	return append(c, element)
}
