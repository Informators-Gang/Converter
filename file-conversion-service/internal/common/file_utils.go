package common

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const uploadPath = "./tmp/file-converter/uploads"
const maxUploadSize = 10 << 20 // 10 MB
const fileRetentionDurationInHours = 4
const fileRetentionDuration = fileRetentionDurationInHours * time.Hour

func isAllowedFileType(fileName string) bool {
    allowedFileTypes := map[string]bool{
        "image/jpeg": true,
        "image/png":  true,
        "image/gif":  true,
    }

    ext := strings.ToLower(filepath.Ext(fileName))
    mimeType := mime.TypeByExtension(ext)

    return allowedFileTypes[mimeType]
}

func startFileCleaner() {
    for {
        time.Sleep(fileRetentionDuration)
        deleteOldFiles()
    }
}

func deleteOldFiles() {
    files, err := os.ReadDir(uploadPath)
    if err != nil {
        fmt.Println("Error reading directory:", err)
        return
    }

    for _, file := range files {
        filePath := filepath.Join(uploadPath, file.Name())
        fileInfo, err := os.Stat(filePath)
        if err != nil {
            fmt.Println("Error stating file:", err)
            continue
        }

        if time.Since(fileInfo.ModTime()) > fileRetentionDuration {
            os.Remove(filePath)
            fmt.Println("Deleted old file:", filePath)
        }
    }
}