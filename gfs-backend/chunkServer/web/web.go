package chunkWeb

import (
	"fmt"
	"gfs-go/pb"
	"log"
	"net"

	chunkApp "gfs-go/chunkServer/application"
	chunkDomain "gfs-go/chunkServer/domain"

	"google.golang.org/grpc"
)

var server *chunkDomain.ChunkServer

func createServer(ready chan<- struct{}) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	server.Addr = lis.Addr().String()
	s := grpc.NewServer()
	pb.RegisterChunkServerServiceServer(s, server)
	fmt.Println("localhost:" + server.Addr)
	close(ready)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Fatal to serve")
	}
}
func Start() {
	server = &chunkDomain.ChunkServer{Addr: "", ChunkIds: []string{}, MemoryUtilization: 0}
	ready := make(chan struct{})
	// Forces the start of the heartbeat sending only after the server has been creted
	go func() {
		<-ready
		go chunkApp.SendPeriodicHeartBeat(server)
	}()
	createServer(ready)
}
