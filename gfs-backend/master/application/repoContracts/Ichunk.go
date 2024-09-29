package masterApplicationContracts

type ChunkRepository interface {
	SaveChunk(chunk []byte, fileName string) error
}
