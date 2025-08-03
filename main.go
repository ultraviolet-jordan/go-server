package main

import (
	"awesomeProject/cache"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// ---- load environment properties
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	members, err := strconv.ParseBool(os.Getenv("NODE_MEMBERS"))
	if err != nil {
		log.Fatalf("Invalid NODE_MEMBERS value: %v", err)
	}
	// ---- load cache
	start := time.Now()
	cache.LoadObjs(members, "./data/pack")
	elapsed := time.Since(start)
	log.Printf("Cache loading took %s", elapsed)
}
