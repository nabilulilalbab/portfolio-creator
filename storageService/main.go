package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/nabilulilalbab/library"
	"github.com/nabilulilalbab/storageService/handler"
)

//go:embed static templates/*.html
var appFS embed.FS

func main() {
	// Working directory
	wd, _ := os.Getwd()
	log.Println("Working directory:", wd)

	// Inisialisasi template dari embedded filesystem
	templates, err := handler.InitTemplates(appFS)
	if err != nil {
		log.Fatalf("Error initializing templates: %v", err)
	}

	mux := http.NewServeMux()
	
	// Serve static files
	fileserver := http.FileServer(http.FS(appFS))
	mux.Handle("/static/", library.FileServerWithLog(fileserver))
	
	// Handle routes with error handling middleware
	mux.HandleFunc("/", library.AppHandler(handler.IndexHandler).CreateHandler(library.ErrorHandlerOptions{
		Templates: templates,
	}))

	mux.HandleFunc("/error", library.AppHandler(handler.ErrorSimulasi).CreateHandler(library.ErrorHandlerOptions{
		Templates: templates,
	}))

	// Start server with proper error handling
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
