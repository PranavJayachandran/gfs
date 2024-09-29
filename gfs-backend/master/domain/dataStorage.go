package serverDomain

type ChunkInfo struct {
	ChunkName  string
	ByteOffset int
}
type ChunkToChunkServer struct {
	ChunkName       string
	ChunkServerAddr []string
}

// Will store file to chunk split ups
var FileToChunkMapper map[string][]ChunkInfo

// This will stor the chunk to restApi map for each chunk
var ChunkToChunkServerMapper map[string][]string
