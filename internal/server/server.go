package server

import (
	"fmt"
	i "github.com/thisismeamir/kage/internal/init"
	configAPI "github.com/thisismeamir/kage/internal/server/api/v1/config"
	configAtoms "github.com/thisismeamir/kage/internal/server/api/v1/config/atoms"
	configModules "github.com/thisismeamir/kage/internal/server/api/v1/config/modules"

	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

////go:embed ../../web/out/*
//var staticFiles embed.FS

type Server struct {
	router *gin.Engine
}

func New() *Server {
	router := gin.Default()

	addr := fmt.Sprintf("http://%s:%d", i.GetGlobalConfig().Client.Web.Host, i.GetGlobalConfig().Client.Web.Port)
	log.Println("Answering CORS requests from:", addr)
	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{addr, "*"}, // Next.js dev server
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//// Serve embedded Next.js static files
	//router.StaticFS("/static", http.FS(staticFiles))
	//router.GET("/", func(c *gin.Context) {
	//	c.FileFromFS("web/out/index.html", http.FS(staticFiles))
	//})

	// API routes
	api := router.Group("/api/v1")
	{
		// Get configuration file
		api.GET("/config", configAPI.GetConfiguration)
		// Atom Paths
		api.GET("/config/atoms", configAtoms.GetAtomPaths)      // Get all atom paths
		api.POST("/config/atoms", configAtoms.AddAtomPath)      // Add a new atom path
		api.DELETE("/config/atoms", configAtoms.DeleteAtomPath) // Add a new atom path
		// Module Paths
		api.GET("/config/modules", configModules.GetModulePaths)      // Get all atom paths
		api.POST("/config/modules", configModules.AddModulePath)      // Add a new atom path
		api.DELETE("/config/modules", configModules.DeleteModulePath) // Add a new atom path
	}

	return &Server{router: router}
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func addPathOfAtoms(c *gin.Context) {
	var req config.AtomPathRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	// Example response (replace with your logic)
	resp := config.AtomPathResponse{
		AtomPath: req.Path,
		Exists:   true, // or your actual check
		Valid:    true, // or your actual validation
	}
	c.JSON(200, resp)
}
