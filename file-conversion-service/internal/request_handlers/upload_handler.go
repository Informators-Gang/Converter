package request_handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Informators-Gang/Converter/file-conversion-service/internal/custom_common"
	"github.com/google/uuid"
)

type FileUploadingResponse struct {
	FileID  string `json:"file_id"`
}

func FileUploadingHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    r.ParseMultipartForm(custom_common.MAX_UPLOAD_SIZE)
    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Invalid file", http.StatusBadRequest)
		log.Print("Failed uploading file. Invalid file.")
        return
    }
    defer file.Close()

    if !custom_common.IsAllowedFileType(header.Filename) {
        http.Error(w, "File type not allowed", http.StatusBadRequest)
		log.Print("Failed uploading file. File type not allowed.")
        return
    }

    uniqueID := uuid.New().String()
    originalFileName := header.Filename
    newFileName := uniqueID + "_" + filepath.Clean(originalFileName)
    filePath := filepath.Join(custom_common.UPLOAD_PATH, newFileName)

    dst, err := os.Create(filePath)
    if err != nil {
        http.Error(w, "Failed to save file", http.StatusInternalServerError)
		log.Printf("Failed to create file: %s", err)
        return
    }
    defer dst.Close()

    _, err = io.Copy(dst, file)
    if err != nil {
        http.Error(w, "Failed to save file", http.StatusInternalServerError)
		log.Printf("Failed to copy file: %s", err)		
        return
    }
	
	res := FileUploadingResponse{
		FileID:  uniqueID,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	
	log.Printf("File uploaded successfully with ID: %s", uniqueID)
}