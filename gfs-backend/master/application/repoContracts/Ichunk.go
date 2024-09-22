package masterApplicationContracts

type ChunkRepository interface {
	SaveChunk(chunk []byte) error
}
