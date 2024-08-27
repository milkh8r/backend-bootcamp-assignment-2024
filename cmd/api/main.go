package main

import "log"

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	db, err := postgres.NewC
}
