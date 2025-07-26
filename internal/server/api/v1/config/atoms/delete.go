package atoms

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/init"
	"github.com/thisismeamir/kage/internal/models"
	"github.com/thisismeamir/kage/internal/server/config"
)

func DeleteAtomPath(c *gin.Context) {
	var req models.AtomPath
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	atomPaths := i.GetGlobalConfig().AtomPaths
	newAtomPaths := []models.AtomPath{}
	deleted := false
	for _, atomPath := range atomPaths {
		if atomPath.Path != req.Path {
			newAtomPaths = append(newAtomPaths, atomPath)
		} else {
			deleted = true
		}
	}

	if deleted {
		cfg := i.GetGlobalConfig()
		cfg.AtomPaths = newAtomPaths
		i.SetGlobalConfig(cfg)
		i.SaveConfigFile()
		resp := config.DeleteAtomPathResponse{
			AtomPath: req.Path,
			Deleted:  true,
			Message:  "Atom path deleted successfully.",
		}
		c.JSON(200, resp)
	} else {
		c.JSON(404, gin.H{"error": "atom path not found"})
	}
}

func deleteModulePath(c *gin.Context) {
	var req models.AtomPath
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	atomPaths := i.GetGlobalConfig().AtomPaths
	newAtomPaths := []models.AtomPath{}
	deleted := false
	for _, atomPath := range atomPaths {
		if atomPath.Path != req.Path {
			newAtomPaths = append(newAtomPaths, atomPath)
		} else {
			deleted = true
		}
	}

	if deleted {
		cfg := i.GetGlobalConfig()
		cfg.AtomPaths = newAtomPaths
		i.SetGlobalConfig(cfg)
		i.SaveConfigFile()
		resp := config.DeleteAtomPathResponse{
			AtomPath: req.Path,
			Deleted:  true,
			Message:  "Atom path deleted successfully.",
		}
		c.JSON(200, resp)
	} else {
		c.JSON(404, gin.H{"error": "atom path not found"})
	}
}
