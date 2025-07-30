package main

import (
	"fmt"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	config2 "github.com/thisismeamir/kage/internal/bootstrap/config"
	. "github.com/thisismeamir/kage/internal/runtime/handler/file"
	. "github.com/thisismeamir/kage/internal/server"
	"log"
	"os"
)

func main() {
	art, _ := os.ReadFile("./util/asci-art")
	fmt.Println(string(art))
	// Initialization Function:
	config := config2.LoadConfiguration(config2.GetConfigPath())
	config2.SetGlobalConfig(config)
	serverAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	serverAddr, config := i.InitKage()

	// Start the server
	log.Println("Starting Server in", serverAddr)
	srv := New()
	if err := srv.Start(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
