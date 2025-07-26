package modules

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/init"
	"github.com/thisismeamir/kage/internal/models"
	"github.com/thisismeamir/kage/internal/server/config"
)

func DeleteModulePath(c *gin.Context) {
	var req models.ModulePath
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	modulePaths := i.GetGlobalConfig().ModulePaths
	newModulePaths := []models.ModulePath{}
	deleted := false
	for _, modulePath := range modulePaths {
		if modulePath.Path != req.Path {
			newModulePaths = append(newModulePaths, modulePath)
		} else {
			deleted = true
		}
	}

	if deleted {
		cfg := i.GetGlobalConfig()
		cfg.ModulePaths = newModulePaths
		i.SetGlobalConfig(cfg)
		i.SaveConfigFile()
		resp := config.DeleteModulePathResponse{
			ModulePath: req.Path,
			Deleted:    true,
			Message:    "Module path deleted successfully.",
		}
		c.JSON(200, resp)
	} else {
		c.JSON(404, gin.H{"error": "Module path not found"})
	}
}
