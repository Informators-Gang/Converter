package request_handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Informators-Gang/Converter/file-conversion-service/internal/custom_common"
)

type ConvertFileResponse struct {
    NewFileID  string `json:"new_file_id"`
}

func ConversionHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        log.Println("Error: Only POST method is allowed")
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    fileID := r.URL.Query().Get("file_id")
    convertTo := r.URL.Query().Get("convert_to")
    if fileID == "" || convertTo == "" {
        log.Println("Error: File ID and conversion format are required")
        http.Error(w, "File ID and conversion format are required", http.StatusBadRequest)
        return
    }

    log.Printf("Converting file with ID %s to format %s\n", fileID, convertTo)

    newFileID, conversionErr := custom_common.ConvertFile(fileID, convertTo)
    if conversionErr != nil {
        log.Println("Error converting file:", conversionErr)
        http.Error(w, conversionErr.Error(), http.StatusInternalServerError)
        return
    }

    log.Printf("File converted successfully, new file ID: %s\n", newFileID)

    response := ConvertFileResponse{
        NewFileID: newFileID,
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Println("Error encoding JSON:", err)
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
    }
}
