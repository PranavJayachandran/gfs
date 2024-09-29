package masterInfrastructure

import (
	"context"
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
	chunkSize := 4
	for i := 0; i < len(chunk); i += chunkSize {
		end := i + chunkSize
		if end > len(chunk) {
			end = len(chunk) // Adjust to the end of the data
		}
		chunkName := getUniqueFileName()
		for range serverDomain.ReplicationFactor {
			sendToChunkServer(i, chunk[i:end], getChunkServerAddr(), fileName, chunkName)
		}

	}
	return nil
}
func getChunkServerAddr() string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(len(serverDomain.MasterServer.ChunkServers))
	return serverDomain.MasterServer.ChunkServers[randomNumber].ServerAddr
}
func sendToChunkServer(offset int, chunk []byte, addr string, fileName string, chunkName string) {
	port := strings.TrimPrefix(addr, "[::]:")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:"+port, opts)
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	client := pb.NewChunkServerServiceClient(cc)
	serverDomain.FileToChunkMapper[fileName] = append(serverDomain.FileToChunkMapper[fileName], serverDomain.ChunkInfo{
		ChunkName:  chunkName,
		ByteOffset: offset,
	})
	serverDomain.ChunkToChunkServerMapper[chunkName] = append(serverDomain.ChunkToChunkServerMapper[chunkName], addr)
	request := &pb.ChunkRequest{
		Chunk:    chunk,
		FileName: chunkName,
	}
	_, err = client.StoreChunk(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
}
