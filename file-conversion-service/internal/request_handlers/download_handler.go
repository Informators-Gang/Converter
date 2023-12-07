package request_handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Informators-Gang/Converter/file-conversion-service/internal/custom_common"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
    // Check if the request method is GET
    if r.Method != http.MethodGet {
        http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
        return
    }

    // Extract file ID from the query parameter
    fileID := r.URL.Query().Get("file_id")
    if fileID == "" {
        http.Error(w, "File ID is required", http.StatusBadRequest)
        return
    }

    // Find the file by ID
    filePath, err := custom_common.FindFileByID(fileID)
    if err != nil {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }

    // Open the file
    file, err := os.Open(filePath)
    if err != nil {
        http.Error(w, "Error opening file", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    // Set response headers
    w.Header().Set("Content-Disposition", "attachment; filename=" + custom_common.RemoveIDFromFileName(filepath.Base(filePath)))
    w.Header().Set("Content-Type", "application/octet-stream")

    // Copy the file content to the response writer
    if _, err := io.Copy(w, file); err != nil {
        http.Error(w, "Error sending file", http.StatusInternalServerError)
    }
}
