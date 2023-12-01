package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Informators-Gang/Converter/file-conversion-service/internal/custom_common"
	"github.com/Informators-Gang/Converter/file-conversion-service/internal/request_handlers"
)

var port = flag.Int("port", 5353, "port to run the server on")

func main() {
    flag.Parse()

    http.HandleFunc("/upload", request_handlers.FileUploadingHandler)
    http.HandleFunc("/convert", request_handlers.ConversionHandler)
    http.HandleFunc("/file_info", request_handlers.FileInfoHandler)
    http.HandleFunc("/download", request_handlers.DownloadHandler)

    if _, err := os.Stat(custom_common.UPLOAD_PATH); os.IsNotExist(err) {
		log.Printf("Folder %s does not exist. Creating...", custom_common.UPLOAD_PATH)
		err := os.MkdirAll(custom_common.UPLOAD_PATH, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

    // Start the file cleaner
    go custom_common.StartFileCleaner()

    addr := fmt.Sprintf(":%d", *port)
    log.Printf("Starting server on %s\n", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}
