package main

import chunkWeb "gfs-go/chunkServer/web"

func main() {
	chunkWeb.Start()

}

//  export PATH="$PATH:$(go env GOPATH)/bin"
// protoc --go_out=. --go-grpc_out=. gfs.proto
