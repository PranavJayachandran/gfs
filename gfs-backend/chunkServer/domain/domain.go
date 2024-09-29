package chunkDomain

import (
	"context"
	"fmt"
	"gfs-go/pb"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ChunkServer struct {
	pb.ChunkServerServiceServer
	RpcAddr           string
	RestAddr          string
	MemoryUtilization float32
	ChunkIds          []string
}

// addr =>  is to be set when the server is spin up
// memoryUtilization => should be updated whenever a chunk is manipulated(add,remove or modified)
// chunksIds => should be updated whenever a chunk is added or removed.

func (s *ChunkServer) StoreChunk(ctx context.Context, req *pb.ChunkRequest) (*emptypb.Empty, error) {
	err := os.MkdirAll("files/"+s.RestAddr, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return &emptypb.Empty{}, nil
	}
	err = os.MkdirAll("files", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return &emptypb.Empty{}, nil
	}
	err = os.WriteFile(filepath.Join("files/"+s.RestAddr, req.FileName), req.Chunk, 0644)
	if err != nil {
		fmt.Printf("File %s couldnot be created\n", req.FileName)
	}
	fmt.Printf("File created: %s", req.FileName)
	return &emptypb.Empty{}, nil
}
