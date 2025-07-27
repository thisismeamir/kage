package atoms

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	"github.com/thisismeamir/kage/pkg/node"
)

// DeleteAtomPathResponse is the response structure for removing an node path.
type DeleteAtomPathResponse struct {
	AtomPath string `json:"atomPath"`
	Deleted  bool   `json:"removed"`
	Message  string `json:"message,omitempty" jsonschema:"omitempty" jsonschema_extras:"description=Message about the removal status"`
}

func DeleteAtomPath(c *gin.Context) {
	var req node.NodePath
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	atomPaths := i.GetGlobalConfig().AtomPaths
	newAtomPaths := []node.NodePath{}
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
		resp := DeleteAtomPathResponse{
			AtomPath: req.Path,
			Deleted:  true,
			Message:  "Atom path deleted successfully.",
		}
		c.JSON(200, resp)
	} else {
		c.JSON(404, gin.H{"error": "node path not found"})
	}
}

func deleteModulePath(c *gin.Context) {
	var req node.NodePath
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	atomPaths := i.GetGlobalConfig().AtomPaths
	newAtomPaths := []node.NodePath{}
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
		resp := DeleteAtomPathResponse{
			AtomPath: req.Path,
			Deleted:  true,
			Message:  "Atom path deleted successfully.",
		}
		c.JSON(200, resp)
	} else {
		c.JSON(404, gin.H{"error": "node path not found"})
	}
}
