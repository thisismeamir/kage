package main

import (
	"fmt"
	"github.com/thisismeamir/kage/internal/bootstrap/init_methods"
	"log"
	"os"
)

var serverAddr string
var serverPort string
var apiPath string
var apiVersion string

func init() {
	// ASCII ART FOR RUNNING THE PROGRAM IN CLI.
	art, _ := os.ReadFile("./util/asci-art")
	fmt.Println(string(art))
	// Loading Configuration
	config, err := init_methods.LoadConfiguration("/home/kid-a/projects/kage/configs/default.conf.json")
	if err == nil {
		err := init_methods.SetupBasePath(config)
		if err != nil {
			log.Fatalf("[FATAL] Unable to set up base path, and directories. %s", err)
		}
	} else {
		log.Fatalf("[FATAL] Unable to load configuration. %s", err)
	}

	// initializing logger
	init_methods.InitializeLogger(config.BasePath + config.Logging.File)

	files := init_methods.GetPathsObjects(config.Paths)
	for _, file := range files {
		log.Println(file)
		value := init_methods.GetTypeOfObject(file)
		println(value)
	}
}

func main() {

	// Initialization Function:
	//init_methods := config2.LoadConfiguration(config2.GetConfigPath())
	//config2.SetGlobalConfig(init_methods)
	//
	//serverAddr, init_methods, watchers := i.InitKage()
	//
	//for _, watcher := range watchers {
	//	go func() {
	//		_ = watcher.Start()
	//	}()
	//}
	//// Start the server
	//log.Println("Starting Server in", serverAddr)
	//srv := New()
	//if err := srv.Start(serverAddr); err != nil {
	//	log.Fatal("Failed to start server:", err)
	//}
}
