package request_handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Informators-Gang/Converter/file-conversion-service/internal/custom_common"
)

// FileInfoResponse represents the structure of the file information response
type FileInfoResponse struct {
    Filename           string   `json:"filename"`
    Size               int64    `json:"size"`
    Extension          string   `json:"extension"`
    ConvertibleFormats []string `json:"convertible_formats"`
}

func FileInfoHandler(w http.ResponseWriter, r *http.Request) {
    // Check if the request method is GET
    if r.Method != http.MethodGet {
        log.Println("Error: Only GET method is allowed")
        http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
        return
    }

    // Get the file ID from the query string
    fileID := r.URL.Query().Get("file_id")
    if fileID == "" {
        log.Println("Error: File ID is required")
        http.Error(w, "File ID is required", http.StatusBadRequest)
        return
    }

    log.Printf("Retrieving info for file with ID %s\n", fileID)

    // Construct the file path
    filePath, err := custom_common.FindFileByID(fileID)
    if err != nil {
        log.Println("Error finding file by ID:", err)
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    // Get file info
    fileInfo, err := os.Stat(filePath)
    if err != nil {
        log.Println("Error getting file info:", err)
        http.Error(w, "Error getting file info", http.StatusInternalServerError)
        return
    }

    log.Printf("File info retrieved successfully for file with ID %s\n", fileID)

    // Extract extension and determine convertible formats
    extension := strings.TrimPrefix(filepath.Ext(fileInfo.Name()), ".")
    convertibleFormats := custom_common.GetConvertibleFormats(extension)

    // Prepare the file info response
    response := FileInfoResponse{
        Filename:           custom_common.RemoveIDFromFileName(fileInfo.Name()),
        Size:               fileInfo.Size(),
        Extension:          extension,
        ConvertibleFormats: convertibleFormats,
    }

    // Set the Content-Type header
    w.Header().Set("Content-Type", "application/json")

    // Encode and send the response
    err = json.NewEncoder(w).Encode(response)
    if err != nil {
        log.Println("Error encoding JSON:", err)
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
    }
}