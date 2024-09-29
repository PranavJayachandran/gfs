package chunkWeb

import (
	"fmt"
	"gfs-go/pb"
	"log"
	"net"
	"net/http"

	chunkApp "gfs-go/chunkServer/application"
	chunkDomain "gfs-go/chunkServer/domain"
	ChunkController "gfs-go/chunkServer/web/Controller"

	"google.golang.org/grpc"
)

var server *chunkDomain.ChunkServer

func createRpcServer(ready chan<- struct{}) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	server.RpcAddr = lis.Addr().String()
	server.RestAddr = lis.Addr().String()
	s := grpc.NewServer()
	pb.RegisterChunkServerServiceServer(s, server)
	fmt.Println("localhost:" + server.RpcAddr)
	close(ready)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Fatal to serve grpc server")
	}
}
func startRestServer() {
	http.HandleFunc("/getFile", ChunkController.GetFile)
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	server.RestAddr = lis.Addr().String()
	if err := http.Serve(lis, nil); err != nil {
		log.Fatal("Failed to spin up http sever")
	}
}
func Start() {
	server = &chunkDomain.ChunkServer{RpcAddr: "", ChunkIds: []string{}, MemoryUtilization: 0}
	ready := make(chan struct{})
	go startRestServer()
	// Forces the start of the heartbeat sending only after the server has been creted
	go func() {
		<-ready
		go chunkApp.SendPeriodicHeartBeat(server)
	}()
	createRpcServer(ready)
}
