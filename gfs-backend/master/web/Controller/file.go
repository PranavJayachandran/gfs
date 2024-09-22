package masterWebController

import (
	masterApplication "gfs-go/master/application"
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
	file, _, err := r.FormFile("file")
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

	fileProcessor.SendToChunkServers(content)
	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully!"))
}
