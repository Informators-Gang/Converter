package request_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Informators-Gang/Converter/file-conversion-service/internal/custom_common"
)

type ConvertFileResponse struct {
    Message    string `json:"message"`
    NewFileID  string `json:"new_file_id"`
}

func ConversionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

	fileID := r.URL.Query().Get("file_id")
	convertTo := r.URL.Query().Get("convert_to")
	if fileID == "" || convertTo == "" {
		http.Error(w, "File ID and conversion format are required", http.StatusBadRequest)
		return
	}

    newFileID, conversionErr := custom_common.ConvertFile(fileID, convertTo)
    if conversionErr != nil {
        http.Error(w, conversionErr.Error(), http.StatusInternalServerError)
        return
    }

    response := ConvertFileResponse{
        Message:   "File converted successfully",
        NewFileID: newFileID,
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
    }
}
