package serverDomain

type ChunkInfo struct {
	ChunkName  string
	ByteOffset int
}

var FileToChunkMapper map[string][]ChunkInfo
var ChunkToChunkServerMapper map[string][]string
