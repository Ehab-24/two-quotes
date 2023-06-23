package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"suraj.com/refine/routers"
)

const port = 3000

func loadEnv() {
	fileName := ".env.local"
	if err := godotenv.Load(fileName); err != nil {
		log.Fatalf("Error loading %v\nError: %v", fileName, err)
	}
}

func main() {
	loadEnv()

	fmt.Printf("Server starting on port %v...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), routers.GetBaseRouter()))
}
