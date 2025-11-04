package main

import "backend-service/server"
import "github.com/joho/godotenv"
import "log"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not loaded or missing")
	}
	server.NewServer().ListenAndServe()	
}