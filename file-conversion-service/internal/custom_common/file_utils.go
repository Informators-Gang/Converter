package custom_common

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/image/tiff"
)

const UPLOAD_PATH = "/tmp/file-converter/uploads"
const MAX_UPLOAD_SIZE = 10 << 20 // 10 MB
const FILE_RETENTION_DURATION_IN_HOURS = 4
const FILE_RETENTION_DURATION = FILE_RETENTION_DURATION_IN_HOURS * time.Hour

func IsAllowedFileType(fileName string) bool {
    allowedFileTypes := map[string]bool{
        "image/jpeg": true,
        "image/png":  true,
        "image/gif":  true,
    }

    ext := strings.ToLower(filepath.Ext(fileName))
    mimeType := mime.TypeByExtension(ext)

    return allowedFileTypes[mimeType]
}

func StartFileCleaner() {
    deleteOldFiles()
    for {
        time.Sleep(FILE_RETENTION_DURATION)
        deleteOldFiles()
    }
}

func deleteOldFiles() {
    files, err := os.ReadDir(UPLOAD_PATH)
    if err != nil {
        fmt.Println("Error reading directory:", err)
        return
    }

    for _, file := range files {
        filePath := filepath.Join(UPLOAD_PATH, file.Name())
        fileInfo, err := os.Stat(filePath)
        if err != nil {
            fmt.Println("Error stating file:", err)
            continue
        }

        if time.Since(fileInfo.ModTime()) > FILE_RETENTION_DURATION {
            os.Remove(filePath)
            fmt.Println("Deleted old file:", filePath)
        }
    }
}

func FindFileByID(fileID string) (string, error) {
    files, err := os.ReadDir(UPLOAD_PATH)
    if err != nil {
        return "", err
    }

    for _, file := range files {
        if strings.HasPrefix(file.Name(), fileID+"_") {
            return filepath.Join(UPLOAD_PATH, file.Name()), nil
        }
    }

    return "", fmt.Errorf("file not found")
}

func GetConvertibleFormats(extension string) []string {
    switch extension {
    case "jpg", "jpeg":
        return []string{"png", "tiff"}
    case "tiff":
        return []string{"jpg", "png"}
    case "png":
        return []string{"jpg", "tiff"}
    default:
        return []string{}
    }
}

func RemoveIDFromFileName(fileName string) string {
    parts := strings.SplitN(fileName, "_", 2)
    if len(parts) < 2 {
        return fileName
    }
    return parts[1]
}

func ConvertFile(fileID, convertTo string) (string, error) {
    // Find the original file by ID
    originalFilePath, err := FindFileByID(fileID)
    if err != nil {
        return "", err
    }

    // Open the original file
    originalFile, err := os.Open(originalFilePath)
    if err != nil {
        return "", err
    }
    defer originalFile.Close()

    // Decode the original image
    img, _, err := image.Decode(originalFile)
    if err != nil {
        return "", err
    }

    // Generate a new unique ID for the converted file
    newFileID := uuid.New().String()
    newFileName := newFileID + "." + convertTo
    newFilePath := filepath.Join(UPLOAD_PATH, newFileName)

    // Create the new file
    newFile, err := os.Create(newFilePath)
    if err != nil {
        return "", err
    }
    defer newFile.Close()

    // Convert and save the new image
    switch strings.ToLower(convertTo) {
    case "jpeg", "jpg":
        err = jpeg.Encode(newFile, img, nil)
    case "png":
        err = png.Encode(newFile, img)
    case "tiff":
        err = tiff.Encode(newFile, img, nil)
    default:
        return "", fmt.Errorf("unsupported conversion format: %s", convertTo)
    }

    if err != nil {
        return "", err
    }

    return newFileID, nil
}