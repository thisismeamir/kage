package server

import (
	"github.com/thisismeamir/kage/internal/server/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

////go:embed ../../web/dist/static.txt
//var staticFiles embed.FS

type Server struct {
	router *gin.Engine
}

func New() *Server {
	router := gin.Default()

	//// Serve embedded frontend
	//router.StaticFS("/static", http.FS(staticFiles))
	//router.GET("/", func(c *gin.Context) {
	//	c.FileFromFS("web/dist/index.html", http.FS(staticFiles))
	//})

	// API routes
	api := router.Group("/api/v1")
	{
		api.GET("/status", getStatus)
		api.GET("/modules", getModules)
		api.POST("/modules", createModule)
		api.GET("/atoms", getAtoms)
		api.POST("/atoms", createAtom)
	}

	return &Server{router: router}
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

// Placeholder handlers
func getStatus(c *gin.Context) {
	c.JSON(200, gin.H{"status": "running", "modules": 0})
}

func getModules(c *gin.Context) {
	c.JSON(200, []interface{}{})
}

func createModule(c *gin.Context) {
	c.JSON(201, gin.H{"message": "module created"})
}

func getAtoms(c *gin.Context) {
	c.JSON(200, []interface{}{})
}

func createAtom(c *gin.Context) {
	c.JSON(201, gin.H{"message": "atom created"})
}
