package ChunkController

import (
	"fmt"
	chunkDomain "gfs-go/chunkServer/domain"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func GetFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fileName := r.URL.Query().Get("fileName")
	if fileName == "" {
		http.Error(w, "FileName is required", http.StatusBadRequest)
	}
	folderName := chunkDomain.Server.RestAddr + chunkDomain.Server.RpcAddr
	filePath := filepath.Join("files/" + folderName + "/" + fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found"+filePath, http.StatusNotFound)
		return
	}

	// Read the file contents
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Set the content type based on the file type (optional, plain text for simplicity)
	w.Header().Set("Content-Type", "text/plain")

	// Write the file contents as the response
	w.Write(content)
}
