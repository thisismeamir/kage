package main

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/thisismeamir/kage/internal/bootstrap/init_methods"
	engine2 "github.com/thisismeamir/kage/internal/engine"
	execution_system "github.com/thisismeamir/kage/internal/engine/execution-system"
	system_monitor "github.com/thisismeamir/kage/internal/engine/system-monitor"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	"github.com/thisismeamir/kage/internal/server"
	"github.com/thisismeamir/kage/internal/watcher"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
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
	// Load registry and create engine components
	var err error
	reg, err = registry.LoadRegistry(conf.BasePath + "/data/registry.json")
	if err != nil {
		log.Fatalf("[FATAL] Unable to load registry: %s", err)
	}

	sm := system_monitor.NewSystemMonitor()
	ex := execution_system.NewExecutionSystem()

	// Create the engine
	engine := engine2.NewEngine(conf, *reg, *ex, *sm)

	// Create the server
	srv := server.New(clientAddr)

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Channel to listen for interrupt signal for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start the engine routine in a separate goroutine
	wg.Add(1)
	go func() {
		log.Println("Starting engine routine...")
		engine.Routine(ctx, &wg)
	}()

	// Start the server in a separate goroutine with timeout handling
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("Starting server on %s...", serverAddr)

		// Create a channel to receive server start result
		serverErr := make(chan error, 1)

		// Start server in another goroutine
		go func() {
			serverErr <- srv.Start(serverAddr)
		}()

		// Wait for either context cancellation or server error
		select {
		case <-ctx.Done():
			log.Println("Server stopping due to context cancellation...")
			return
		case err := <-serverErr:
			if err != nil {
				log.Printf("Server error: %v", err)
				cancel() // Cancel context to stop other components
			}
		}
	}()

	// Wait for interrupt signal
	go func() {
		<-sigChan
		log.Println("Received interrupt signal, shutting down gracefully...")
		cancel() // This will trigger both engine and server to stop

		// Give components time to shutdown gracefully
		time.Sleep(2 * time.Second)
	}()

	// Wait for all goroutines to complete
	log.Println("Application started successfully. Press Ctrl+C to stop.")
	wg.Wait()
	log.Println("Application stopped gracefully.")
}
