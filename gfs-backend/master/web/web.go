package masterWeb

import (
	"fmt"
	constant "gfs-go/constants"
	serverDomain "gfs-go/master/domain"
	masterWebController "gfs-go/master/web/Controller"
	"gfs-go/pb"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

func Start() {
	go startGRpcServer()
	startEndpointServer()
}
func startEndpointServer() {
	http.HandleFunc("/uploadFile", masterWebController.UploadFile)
	http.HandleFunc("/getFile", masterWebController.GetFileData)
	fmt.Println("Http server started at " + constant.MASTER_ENDPOINT)
	if err := http.ListenAndServe(":"+constant.MASTER_ENDPOINT, nil); err != nil {
		log.Fatalf("Failed to serve the endpoint")
	}
}
func startGRpcServer() {
	serverDomain.FileToChunkMapper = make(map[string][]serverDomain.ChunkInfo)
	serverDomain.ChunkToChunkServerMapper = make(map[string][]string)
	lis, err := net.Listen("tcp", ":"+constant.MASTER_ADDR_PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterMasterServiceServer(s, &serverDomain.Server{})
	fmt.Println("GRPC started at" + constant.MASTER_ADDR_PORT)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fatal to serve")
	}
}
