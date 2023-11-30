package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Informators-Gang/Converter/file-conversion-service/internal/request_handlers"
)

var port = flag.Int("port", 8000, "port to run the server on")

func main() {
    flag.Parse()

    http.HandleFunc("/upload", request_handlers.HandleFileUpload)
    http.HandleFunc("/convert", request_handlers.HandleConversionRequest)

    addr := fmt.Sprintf(":%d", *port)
    log.Printf("Starting server on %s\n", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}
