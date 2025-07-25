package modules

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	"github.com/thisismeamir/kage/pkg/module"
)

// DeleteModulePathResponse is the response structure for removing an atom path.
type DeleteModulePathResponse struct {
	ModulePath string `json:"module_path"`
	Deleted    bool   `json:"removed"`
	Message    string `json:"message,omitempty" jsonschema:"omitempty" jsonschema_extras:"description=Message about the removal status"`
}

func DeleteModulePath(c *gin.Context) {
	var req atom.ModulePath
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	modulePaths := i.GetGlobalConfig().ModulePaths
	newModulePaths := []atom.ModulePath{}
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
		resp := DeleteModulePathResponse{
			ModulePath: req.Path,
			Deleted:    true,
			Message:    "Module path deleted successfully.",
		}
		c.JSON(200, resp)
	} else {
		c.JSON(404, gin.H{"error": "Module path not found"})
	}
}
