// This pacakge would be responsible to taking the channel inputs and make the required functions run

package chunkApp

import (
	"context"
	"fmt"
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
	fmt.Println("Sending heartbeat")
	request := &pb.HeartBeatRequest{
		Addr:              server.Addr,
		ChunkIds:          server.ChunkIds,
		MemoryUtilization: float32(server.MemoryUtilization),
	}
	res, err := client.HeartBeat(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RE")
	fmt.Println(res)
}
func SendPeriodicHeartBeat(server *chunkDomain.ChunkServer) {
	for {
		sendHeartBeatMessage(server)
		time.Sleep(5 * time.Second)
	}
}
