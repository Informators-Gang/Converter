package request_handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Informators-Gang/Converter/file-conversion-service/internal/custom_common"
	"github.com/google/uuid"
)

type Response struct {
	FileID  string `json:"file_id"`
}

func HandleFileUpload(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    r.ParseMultipartForm(custom_common.MAX_UPLOAD_SIZE)
    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Invalid file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    if !custom_common.IsAllowedFileType(header.Filename) {
        http.Error(w, "File type not allowed", http.StatusBadRequest)
        return
    }

    uniqueID := uuid.New().String()
    originalFileName := header.Filename
    newFileName := uniqueID + "_" + filepath.Clean(originalFileName)
    filePath := filepath.Join(custom_common.UPLOAD_PATH, newFileName)

    dst, err := os.Create(filePath)
    if err != nil {
        http.Error(w, "Failed to save file", http.StatusInternalServerError)
        return
    }
    defer dst.Close()

    _, err = io.Copy(dst, file)
    if err != nil {
        http.Error(w, "Failed to save file", http.StatusInternalServerError)
        return
    }
	
	res := Response{
		FileID:  uniqueID,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	
	fmt.Fprintf(w, "File uploaded successfully with ID: %s", uniqueID)
}