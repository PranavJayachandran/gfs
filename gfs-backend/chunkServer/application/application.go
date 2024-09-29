// This pacakge would be responsible to taking the channel inputs and make the required functions run

package chunkApp

import (
	"context"
	chunkDomain "gfs-go/chunkServer/domain"
	constant "gfs-go/constants"
	"gfs-go/pb"
	"log"
	"time"

	"google.golang.org/grpc"
)

func sendHeartBeatMessage(server *chunkDomain.ChunkServer) {

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:"+constant.MASTER_ADDR_PORT, opts)
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	client := pb.NewMasterServiceClient(cc)
	request := &pb.HeartBeatRequest{
		Rpcaddr:           server.RpcAddr,
		Restaddr:          server.RestAddr,
		ChunkIds:          server.ChunkIds,
		MemoryUtilization: float32(server.MemoryUtilization),
	}
	_, err = client.HeartBeat(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
}
func SendPeriodicHeartBeat(server *chunkDomain.ChunkServer) {
	for {
		sendHeartBeatMessage(server)
		time.Sleep(5 * time.Second)
	}
}
