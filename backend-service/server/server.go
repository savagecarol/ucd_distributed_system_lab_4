package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)


func NewServer() *http.Server {
	log.Print("starting server...")
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {	
		port = 8080
		log.Print("using default port 8080")
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return server
}
