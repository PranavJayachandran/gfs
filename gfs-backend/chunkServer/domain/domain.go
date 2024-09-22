package chunkDomain

import (
	"context"
	"fmt"
	"gfs-go/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ChunkServer struct {
	pb.ChunkServerServiceServer
	Addr              string
	MemoryUtilization float32
	ChunkIds          []string
}

// addr =>  is to be set when the server is spin up
// memoryUtilization => should be updated whenever a chunk is manipulated(add,remove or modified)
// chunksIds => should be updated whenever a chunk is added or removed.

func (s *ChunkServer) StoreChunk(ctx context.Context, req *pb.ChunkRequest) (*emptypb.Empty, error) {
	fmt.Println(req)
	return &emptypb.Empty{}, nil
}
