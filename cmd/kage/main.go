package main

import (
	"fmt"
	initialization "github.com/thisismeamir/kage/internal/init"
	"github.com/thisismeamir/kage/internal/server"
	"log"
	"os"
)

func main() {

	config := initialization.LoadConfiguration("./configs/default.conf.json")
	serverAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	art, _ := os.ReadFile("./util/asci-art")
	fmt.Println(string(art))
	log.Println("Starting Server in", serverAddr)
	srv := server.New(config)
	if err := srv.Start(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
