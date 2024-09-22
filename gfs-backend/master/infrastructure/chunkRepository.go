package masterInfrastructure

import (
	"context"
	"fmt"
	serverDomain "gfs-go/master/domain"
	"gfs-go/pb"
	"log"
	"strings"

	"google.golang.org/grpc"
)

type ChunkRepository struct{}

func NewChunkRepository() *ChunkRepository {
	return &ChunkRepository{}
}
func (c *ChunkRepository) SaveChunk(chunk []byte) error {
	chunkSize := 4
	for i := 0; i < len(chunk); i += chunkSize {
		end := i + chunkSize
		if end > len(chunk) {
			end = len(chunk) // Adjust to the end of the data
		}
		sendToChunkServer(chunk[i:end], getChunkServerAddr())
	}

	fmt.Printf("Received file content:\n%s\n", chunk)
	return nil
}
func getChunkServerAddr() string {
	if len(serverDomain.MasterServer.ChunkServers) > 0 {
		return serverDomain.MasterServer.ChunkServers[0].ServerAddr
	}
	return ""
}
func sendToChunkServer(chunk []byte, addr string) {
	port := strings.TrimPrefix(addr, "[::]:")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:"+port, opts)
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	client := pb.NewChunkServerServiceClient(cc)
	fmt.Println("Sending Chunk")
	request := &pb.ChunkRequest{
		Chunk:    chunk,
		FileName: "x",
	}
	res, err := client.StoreChunk(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
