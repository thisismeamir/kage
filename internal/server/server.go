package server

import (
	"github.com/thisismeamir/kage/internal/server/config"
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

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Next.js dev server
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
		api.POST("/config/atoms_path", addPathOfAtoms)
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
