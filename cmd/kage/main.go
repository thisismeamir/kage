package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/thisismeamir/kage/internal/bootstrap/init_methods"
	"github.com/thisismeamir/kage/internal/server"
	"github.com/thisismeamir/kage/internal/watcher"
	"log"
	"os"
)

var serverAddr string
var clientAddr string
var api string

func init() {
	// ASCII ART FOR RUNNING THE PROGRAM IN CLI.
	art, _ := os.ReadFile("./util/asci-art")
	fmt.Println(string(art))
	// Loading Configuration
	config, err := init_methods.LoadConfiguration("/home/kid-a/projects/kage/configs/default.conf.json")
	if err == nil {
		// Setting up (making sure) paths inside the base path, e.g. data, logs, cache, tmp, etc.
		err := init_methods.SetupBasePath(config)
		if err != nil {
			log.Fatalf("[FATAL] Unable to set up base path, and directories. %s", err)
		}
	} else {
		log.Fatalf("[FATAL] Unable to load configuration. %s", err)
	}

	// Watch for changes in config.Paths and update registry on change
	w, err := watcher.NewWatcher(config.Paths, func(event watcher.FileSystemEvent) {
		if event.Event == fsnotify.Create || event.Event == fsnotify.Remove {
			initRegErr := init_methods.InitializeRegistries(config.Paths, config.BasePath+"/data/registry.json")
			if initRegErr != nil {
				log.Printf("Failed to update registry: %v", initRegErr)
			} else {
				log.Printf("Registry updated due to file event: %s", event.Path)
			}
		}

	})
	if err != nil {
		log.Fatalf("[FATAL] Unable to start watcher: %v", err)
	}
	if err := w.Start(); err != nil {
		log.Fatalf("[FATAL] Watcher failed to start: %v", err)
	}

	// setting up global values:
	serverAddr = fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	clientAddr = fmt.Sprintf("http://%s:%d", config.Client.Web.Host, config.Client.Web.Port)

	println(clientAddr)
	api = fmt.Sprintf("%s/%s", config.Server.Api.BaseUrl, config.Server.Api.Version)
	// initializing logger
	init_methods.InitializeLogger(config.BasePath + config.Logging.File)

	// registry initializer
	initRegErr := init_methods.InitializeRegistries(config.Paths, config.BasePath+"/data/registry.json")
	if initRegErr != nil {
		log.Fatalf("[FATAL] Unable to initialize registries. %s", initRegErr)
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
	// }
	//Start the server
	log.Println("Starting Server in", serverAddr)
	srv := server.New(clientAddr)
	if err := srv.Start(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
