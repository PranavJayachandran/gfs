package masterWebController

import (
	"encoding/json"
	"fmt"
	masterApplication "gfs-go/master/application"
	serverDomain "gfs-go/master/domain"
	masterInfrastructure "gfs-go/master/infrastructure"
	"io/ioutil"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from the request
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	chunkRepo := masterInfrastructure.NewChunkRepository()
	fileProcessor := masterApplication.NewFileProcessor(chunkRepo)

	fileProcessor.SendToChunkServers(content, header.Filename)
	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully!"))
}
func GetFileData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query()
	fileName := query.Get("fileName")
	var ChunkData = serverDomain.FileToChunkMapper[fileName]
	ChunkToChunkServerData := make([]serverDomain.ChunkToChunkServer, 0)
	for _, item := range ChunkData {
		ChunkToChunkServerData = append(ChunkToChunkServerData, serverDomain.ChunkToChunkServer{
			ChunkServerAddr: serverDomain.ChunkToChunkServerMapper[item.ChunkName],
			ChunkName:       item.ChunkName,
		})
	}

	fmt.Println(serverDomain.FileToChunkMapper[fileName])
	w.Header().Set("Content-Type", "application/json")

	// Convert the array of objects into JSON
	err := json.NewEncoder(w).Encode(ChunkToChunkServerData)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
