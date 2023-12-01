package request_handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Informators-Gang/Converter/file-conversion-service/internal/custom_common"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
    // Check if the request method is GET
    if r.Method != http.MethodGet {
        log.Println("Error: Only GET method is allowed")
        http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
        return
    }

    // Extract file ID from the query parameter
    fileID := r.URL.Query().Get("file_id")
    if fileID == "" {
        log.Println("Error: File ID is required")
        http.Error(w, "File ID is required", http.StatusBadRequest)
        return
    }

    log.Printf("Downloading file with ID %s\n", fileID)

    // Find the file by ID
    filePath, err := custom_common.FindFileByID(fileID)
    if err != nil {
        log.Println("Error finding file by ID:", err)
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }

    // Open the file
    file, err := os.Open(filePath)
    if err != nil {
        log.Println("Error opening file:", err)
        http.Error(w, "Error opening file", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    // Set response headers
    w.Header().Set("Content-Disposition", "attachment; filename=" + custom_common.RemoveIDFromFileName(filepath.Base(filePath)))
    w.Header().Set("Content-Type", "application/octet-stream")

    // Copy the file content to the response writer
    if _, err := io.Copy(w, file); err != nil {
        log.Println("Error sending file:", err)
        http.Error(w, "Error sending file", http.StatusInternalServerError)
    }

    log.Printf("File with ID %s downloaded successfully\n", fileID)
}
