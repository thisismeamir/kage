package main

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/thisismeamir/kage/internal/bootstrap/init_methods"
	engine2 "github.com/thisismeamir/kage/internal/engine"
	execution_system "github.com/thisismeamir/kage/internal/engine/execution-system"
	system_monitor "github.com/thisismeamir/kage/internal/engine/system-monitor"
	task_manager "github.com/thisismeamir/kage/internal/engine/task-manager"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	"github.com/thisismeamir/kage/internal/watcher"
	"github.com/thisismeamir/kage/util"
	"log"
	"os"
	"sync"
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
	sm := system_monitor.NewSystemMonitor()
	ex := execution_system.NewExecutionSystem()
	newEvent := task_manager.Event{
		Identifier:      task_manager.IdentifierGeneration("event"),
		GraphIdentifier: "Sample-Graph.0.0.1.graph",
		Urgency:         1,
		Input:           util.LoadJson("/home/kid-a/kage/first/sample_input.json"),
	}

	newEvent.ScheduleFlow(conf, *reg)
	// Create a new Engine
	engine := engine2.NewEngine(conf, *reg, *ex, *sm)

	// Create context and wait group for concurrency control
	ctx, _ := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go engine.Routine(ctx, &wg) // Run the engine's routine in the background

	//// For example, if you want to stop it after 30 seconds:
	//time.Sleep(30 * time.Second)
	//cancel() // This will stop the engine

	// Wait for the goroutine to complete before shutting down
	wg.Wait()
	log.Println("Engine stopped")

}
