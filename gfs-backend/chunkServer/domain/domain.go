package chunkDomain

import (
	"context"
	"fmt"
	constant "gfs-go/constants"
	"gfs-go/pb"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChunkServer struct {
	pb.ChunkServerServiceServer
	RpcAddr           string
	RestAddr          string
	MemoryUtilization float32
	ChunkIds          []string
}

var Server *ChunkServer

// addr =>  is to be set when the server is spin up
// memoryUtilization => should be updated whenever a chunk is manipulated(add,remove or modified)
// chunksIds => should be updated whenever a chunk is added or removed.

// The incoming fliename in the request is same as the chunkId
func (s *ChunkServer) StoreChunk(ctx context.Context, req *pb.ChunkRequest) (*emptypb.Empty, error) {
	err := os.MkdirAll("files/"+s.RestAddr+s.RpcAddr, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return &emptypb.Empty{}, nil
	}
	err = os.MkdirAll("files", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return &emptypb.Empty{}, nil
	}
	err = os.WriteFile(filepath.Join("files/"+s.RestAddr+s.RpcAddr, req.FileName), req.Chunk, 0644)
	if err != nil {
		fmt.Printf("File %s couldnot be created\n", req.FileName)
	}
	s.ChunkIds = append(s.ChunkIds, req.FileName)
	fmt.Println(req.FileName)
	s.MemoryUtilization += float32(constant.CHUNK_SIZE) / float32(constant.TOTAL_MEMORY)
	fmt.Println(s.MemoryUtilization)
	return &emptypb.Empty{}, nil
}
func (s *ChunkServer) CopyChunk(ctx context.Context, req *pb.CopyChunkRequest) (*pb.CopyChunkResponse, error) {
	port := strings.TrimPrefix(req.Rpcaddr, "[::]:")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:"+port, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewChunkServerServiceClient(cc)
	request := &pb.ChunkRequest{
		Chunk:    getChunk(req.ChunkId),
		FileName: req.ChunkId,
	}
	fmt.Println(req.ChunkId, req.Rpcaddr)
	_, err = client.StoreChunk(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	//Todo Should delete if req.ShouldDelete is true
	deleteFile(req.ChunkId)
	return &pb.CopyChunkResponse{
		Successful: true,
	}, nil
}
func deleteFile(fileName string) {
	folderName := Server.RestAddr + Server.RpcAddr
	filePath := filepath.Join("files/" + folderName + "/" + fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatal(err)
	}
	e := os.Remove(filePath)
	if e != nil {
		log.Fatal(e)
	}
}
func getChunk(fileName string) []byte {
	folderName := Server.RestAddr + Server.RpcAddr
	filePath := filepath.Join("files/" + folderName + "/" + fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatal(err)
	}

	// Read the file contents
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return content
}
