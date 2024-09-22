package masterApplication

import (
	masterApplicationContracts "gfs-go/master/application/repoContracts"
)

type FileProccessor struct {
	chunkRepo masterApplicationContracts.ChunkRepository
}

func NewFileProcessor(chunkRepo masterApplicationContracts.ChunkRepository) *FileProccessor {
	return &FileProccessor{chunkRepo: chunkRepo}
}

func (fp *FileProccessor) SendToChunkServers(chunk []byte) {
	fp.chunkRepo.SaveChunk(chunk)
}
