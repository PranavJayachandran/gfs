package serverDomain

import (
	"context"
	"fmt"
	"gfs-go/pb"
	"log"
	"sort"
	"strings"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChunkServer struct {
	ServerGrpcAddr    string
	ServerRestAddr    string
	ChunkIds          []string
	MemoryUtilization float32
	Lock              *sync.Mutex
}

type Server struct {
	pb.MasterServiceServer
	ChunkServers []ChunkServer
}

var MasterServer Server
var ReplicationFactor = 2

func (s *Server) HeartBeat(ctx context.Context, req *pb.HeartBeatRequest) (*emptypb.Empty, error) {
	chunkServer := &ChunkServer{
		ServerGrpcAddr:    req.Rpcaddr,
		ChunkIds:          req.ChunkIds,
		MemoryUtilization: req.MemoryUtilization,
		ServerRestAddr:    req.Restaddr,
	}
	s.ChunkServers = addIfNotPresent(s.ChunkServers, *chunkServer)
	MasterServer = *s
	if req.MemoryUtilization == 0 { // New registration
		reallocation(len(s.ChunkServers) - 1)
	}
	fmt.Printf("Heartbeat from %s with utilization %f ", chunkServer.ServerGrpcAddr, chunkServer.MemoryUtilization)
	fmt.Print("ChunkIds")

	for _, chunk := range req.ChunkIds {
		fmt.Print(chunk + " ")
	}
	fmt.Println()
	return &emptypb.Empty{}, nil
}

func addIfNotPresent(c []ChunkServer, element ChunkServer) []ChunkServer {
	for i, v := range c {
		if v.ServerGrpcAddr == element.ServerGrpcAddr {
			c[i] = element
			return c
		}
	}
	return append(c, element)
}
func reallocation(index int) {
	var overloadedChunkServers []ChunkServer = make([]ChunkServer, 0)
	for _, chunk := range MasterServer.ChunkServers {
		if chunk.MemoryUtilization > 0.5 {
			overloadedChunkServers = append(overloadedChunkServers, chunk)
		}
	}
	sort.Slice(overloadedChunkServers, func(i, j int) bool {
		return overloadedChunkServers[i].MemoryUtilization > overloadedChunkServers[j].MemoryUtilization
	})
	var round = 0
	for MasterServer.ChunkServers[index].MemoryUtilization < 0.5 {
		var send = false
		for _, chunk := range overloadedChunkServers {
			if len(chunk.ChunkIds) > round {
				go copyChunk(chunk, MasterServer.ChunkServers[index].ServerGrpcAddr, chunk.ChunkIds[round])
			}
		}
		if !send {
			break
		}
		round++
	}
}

func copyChunk(chunkServer ChunkServer, toAddr string, chunkId string) {
	fromAddr := chunkServer.ServerGrpcAddr
	port := strings.TrimPrefix(fromAddr, "[::]:")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:"+port, opts)
	if err != nil {
		log.Println("Fatal to connect to rpc server of " + fromAddr)
		return

	}
	defer cc.Close()
	defer chunkServer.Lock.Unlock()

	chunkServer.Lock.Lock()
	client := pb.NewChunkServerServiceClient(cc)
	request := &pb.CopyChunkRequest{
		ChunkId:      chunkId,
		Rpcaddr:      toAddr,
		ShouldDelete: true,
	}
	_, err = client.CopyChunk(context.Background(), request)
	if err != nil {
		fmt.Printf("Copy chunk %s from %s to %s failed\n", chunkId, fromAddr, toAddr)
	}
}
