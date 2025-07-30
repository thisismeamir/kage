package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/thisismeamir/kage/internal/bootstrap/config"
	"github.com/thisismeamir/kage/pkg/form"
	"os"
)

// AddModulePathResponse is the response structure for checking the existence and validity of an node path.
type AddModulePathResponse struct {
	ModulePath string `json:"module_path"`
	Added      bool   `json:"added"`
	Message    string `json:"message"`
}

func AddModulePath(c *gin.Context) {
	var req form.FormPath

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	if req.Local {
		// Checking if the path exists in the local computer:
		_, e := os.Stat(req.Path)
		if e != nil {
			c.JSON(400, gin.H{"error": "path does not exist"})
			return
		}
		// Check if the path is already in the config
		for _, modulePath := range config.GetGlobalConfig().GraphPaths {
			if modulePath.Path == req.Path {
				c.JSON(200, gin.H{"error": "Module path already exists"})
				return
			}
		}
		// Add the new node path to the global config
		config.SetGlobalConfig(config.Config{
			Name:       config.GetGlobalConfig().Name,
			BasePath:   config.GetGlobalConfig().BasePath,
			GraphPaths: append(config.GetGlobalConfig().GraphPaths, req),
			NodePaths:  config.GetGlobalConfig().NodePaths,
			Version:    config.GetGlobalConfig().Version,
			Server:     config.GetGlobalConfig().Server,
			Client:     config.GetGlobalConfig().Client,
		})

		config.SaveConfigFile()

	}
	// Example response (replace with your logic)
	resp := AddModulePathResponse{
		ModulePath: req.Path,
		Added:      true,
		Message:    "Module path added successfully.",
	}
	c.JSON(200, resp)
}
