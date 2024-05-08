package main

import (
	"github.com/mmmajder/devops-search-service/startup"
	cfg "github.com/mmmajder/devops-search-service/startup/config"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdin)
	log.SetOutput(os.Stderr)
	log.SetOutput(os.Stdout)
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
