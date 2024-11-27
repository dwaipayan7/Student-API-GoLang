package main

import (
	"fmt"
	"net/http"

	"github.com/dwaipayan7/student-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	// Register routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	fmt.Printf("Server running at %s\n", cfg.HTTPServer.Addr)

	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
