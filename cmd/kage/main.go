package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/thisismeamir/kage/internal/bootstrap/init_methods"
	"github.com/thisismeamir/kage/internal/engine/scheduler"
	task_manager "github.com/thisismeamir/kage/internal/engine/task-manager"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
	"github.com/thisismeamir/kage/internal/internal-pkg/event"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	"github.com/thisismeamir/kage/internal/server"
	"github.com/thisismeamir/kage/internal/watcher"
	"log"
	"os"
	"time"
)

var serverAddr string
var clientAddr string
var api string
var reg *registry.Registry
var conf config.Config

func init() {
	// ASCII ART FOR RUNNING THE PROGRAM IN CLI.
	art, _ := os.ReadFile("./util/asci-art")
	fmt.Println(string(art))
	// Loading Configuration
	var err error
	conf, err = init_methods.LoadConfiguration("/home/kid-a/projects/kage/configs/default.conf.json")
	if err == nil {
		// Setting up (making sure) paths inside the base path, e.g. data, logs, cache, tmp, etc.
		err := init_methods.SetupBasePath(conf)
		if err != nil {
			log.Fatalf("[FATAL] Unable to set up base path, and directories. %s", err)
		}
	} else {
		log.Fatalf("[FATAL] Unable to load configuration. %s", err)
	}

	// Initialize the registry
	//reg, err = registry.LoadRegistry(config.BasePath + "/data/registry.json")

	// Watch for changes in config.Paths and update registry on change
	w, err := watcher.NewWatcher(conf.Paths, func(event watcher.FileSystemEvent) {
		if event.Event == fsnotify.Create || event.Event == fsnotify.Remove {
			initRegErr := init_methods.InitializeRegistries(conf.Paths, conf.BasePath+"/data/registry.json")
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
	// Watcher for Task Manager to know about changes in flows directory
	w2, err := watcher.NewWatcher([]string{conf.BasePath + "/tmp/flows"}, func(event watcher.FileSystemEvent) {
		if event.Event == fsnotify.Create || event.Event == fsnotify.Remove {
			log.Println("[INFO] Task Manager detected a change in flows directory, re-initializing Task Manager.")
			task_manager.InitializeTaskManager(conf)
		}

	})
	if err != nil {
		log.Fatalf("[FATAL] Unable to start watcher: %v", err)
	}
	if err := w2.Start(); err != nil {
		log.Fatalf("[FATAL] Watcher failed to start: %v", err)
	}

	// setting up global values:
	serverAddr = fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	clientAddr = fmt.Sprintf("http://%s:%d", conf.Client.Web.Host, conf.Client.Web.Port)

	println(clientAddr)
	api = fmt.Sprintf("%s/%s", conf.Server.Api.BaseUrl, conf.Server.Api.Version)
	// initializing logger
	init_methods.InitializeLogger(conf.BasePath + conf.Logging.File)

	// registry initializer
	regg := init_methods.InitializeRegistries(conf.Paths, conf.BasePath+"/data/registry.json")
	if regg != nil {
		log.Fatalf("[FATAL] Unable to initialize registries. %s", regg)
	}
	endArt, _ := os.ReadFile("./util/main-banner")
	fmt.Println("\n" + string(endArt))
}

func main() {
	reg, _ = registry.LoadRegistry(conf.BasePath + "/data/registry.json")

	tm := task_manager.InitializeTaskManager(conf)
	fmt.Printf("%v ", tm)

	e := event.Event{
		Graph:   "Sample-Graph.0.0.1.graph",
		Date:    time.Now().Format("2006-01-02-15-04-05"),
		Urgency: 2,
	}
	e = e.GenerateIdentifier()
	f := scheduler.Scheduler(e, *reg, conf)
	log.Printf("Scheduler initialized with event: %v\n", f)
	//Start the server
	log.Println("Starting Server in", serverAddr)
	srv := server.New(clientAddr)
	if err := srv.Start(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
