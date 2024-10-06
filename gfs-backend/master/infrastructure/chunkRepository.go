package masterInfrastructure

import (
	"context"
	"fmt"
	constant "gfs-go/constants"
	serverDomain "gfs-go/master/domain"
	"gfs-go/pb"
	"log"
	"math/rand"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type ChunkRepository struct{}

func NewChunkRepository() *ChunkRepository {
	return &ChunkRepository{}
}
func (c *ChunkRepository) SaveChunk(chunk []byte, fileName string) error {
	chunkSize := constant.CHUNK_SIZE
	for i := 0; i < len(chunk); i += chunkSize {
		end := i + chunkSize
		if end > len(chunk) {
			end = len(chunk) // Adjust to the end of the data
		}
		chunkName := getUniqueFileName()
		serverDomain.FileToChunkMapper[fileName] = append(serverDomain.FileToChunkMapper[fileName], serverDomain.ChunkInfo{
			ChunkName:  chunkName,
			ByteOffset: i,
		})
		for range serverDomain.ReplicationFactor {
			go sendToChunkServer(chunk[i:end], getChunkServerAddr(), chunkName)
		}
	}
	return nil
}
func getChunkServerAddr() *serverDomain.ChunkServer {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(len(serverDomain.MasterServer.ChunkServers))
	return &serverDomain.MasterServer.ChunkServers[randomNumber]
}
func sendToChunkServer(chunk []byte, chunkServer *serverDomain.ChunkServer, chunkName string) {
	memoryUtilisationIncrease := float32(constant.CHUNK_SIZE) / float32(constant.TOTAL_MEMORY)
	if chunkServer.MemoryUtilization+memoryUtilisationIncrease > 1 {
		fmt.Println("Couldnt upload to " + chunkServer.ServerGrpcAddr)
		return
	}
	chunkServer.MemoryUtilization += memoryUtilisationIncrease
	port := strings.TrimPrefix(chunkServer.ServerGrpcAddr, "[::]:")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:"+port, opts)
	if err != nil {
		log.Println("Unable to connect to the rpc server of " + port)
	}

	defer cc.Close()

	client := pb.NewChunkServerServiceClient(cc)
	serverDomain.ChunkToChunkServerMapper[chunkName] = append(serverDomain.ChunkToChunkServerMapper[chunkName], *chunkServer)
	request := &pb.ChunkRequest{
		Chunk:    chunk,
		FileName: chunkName,
	}
	_, err = client.StoreChunk(context.Background(), request)
	if err != nil {
		fmt.Printf("Store chunk %s to %s failed\n", chunkName, port)
	}
}
