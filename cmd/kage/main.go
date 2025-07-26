package main

import (
	"fmt"
	i "github.com/thisismeamir/kage/internal/init"
	"github.com/thisismeamir/kage/internal/server"
	"github.com/thisismeamir/kage/internal/watcher"
	"log"
	"os"
)

func main() {

	config := i.LoadConfiguration(i.GetConfigPath())
	i.SetGlobalConfig(config)
	serverAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	art, _ := os.ReadFile("./util/asci-art")
	fmt.Println(string(art))
	log.Println("Starting Server in", serverAddr)
	srv := server.New()
	if err := srv.Start(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
