package server

import (
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

func New(client string) *Server {
	router := gin.Default()

	log.Println("Answering CORS requests from:", client)
	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{client, "*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//// Serve embedded Next.js static file
	//router.StaticFS("/static", http.FS(staticFiles))
	//router.GET("/", func(c *gin.Context) {
	//	c.FileFromFS("web/out/index.html", http.FS(staticFiles))
	//})

	// API routes
	api := router.Group("/api/v1")
	{
		// Get configuration file
		api.GET("/stat", Stat)

	}

	return &Server{router: router}
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func Stat(c *gin.Context) {

}
