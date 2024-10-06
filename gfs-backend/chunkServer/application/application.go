// This pacakge would be responsible to taking the channel inputs and make the required functions run

package chunkApp

import (
	"context"
	"fmt"
	chunkDomain "gfs-go/chunkServer/domain"
	constant "gfs-go/constants"
	"gfs-go/pb"
	"io/ioutil"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

func sendHeartBeatMessage(server *chunkDomain.ChunkServer) {

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:"+constant.MASTER_ADDR_PORT, opts)
	if err != nil {
		log.Println("Unable to connect to the rpc server of " + constant.MASTER_ADDR_PORT)
		return
	}

	defer cc.Close()

	client := pb.NewMasterServiceClient(cc)
	chunkDomain.Server.ChunkIds = getChunkIds()
	chunkDomain.Server.MemoryUtilization = getMemoryUtilization()
	request := &pb.HeartBeatRequest{
		Rpcaddr:           server.RpcAddr,
		Restaddr:          server.RestAddr,
		ChunkIds:          chunkDomain.Server.ChunkIds,
		MemoryUtilization: chunkDomain.Server.MemoryUtilization,
	}
	_, err = client.HeartBeat(context.Background(), request)
	if err != nil {
		log.Println("HeartBeat Failed")
	}
}
func SendPeriodicHeartBeat(server *chunkDomain.ChunkServer) {
	for {
		sendHeartBeatMessage(server)
		time.Sleep(5 * time.Second)
	}
}
func getChunkIds() []string {
	err := os.MkdirAll("files/"+chunkDomain.Server.RestAddr+chunkDomain.Server.RpcAddr, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
	}
	folderName := "files/" + chunkDomain.Server.RestAddr + chunkDomain.Server.RpcAddr
	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		fmt.Println(err)
	}
	var chunkIds = make([]string, 0)
	for _, file := range files {
		chunkIds = append(chunkIds, file.Name())
	}
	return chunkIds
}
func getMemoryUtilization() float32 {
	folderName := "files/" + chunkDomain.Server.RestAddr + chunkDomain.Server.RpcAddr
	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		fmt.Println(err)
	}
	return (float32(len(files) * constant.CHUNK_SIZE)) / float32(constant.TOTAL_MEMORY)
}
