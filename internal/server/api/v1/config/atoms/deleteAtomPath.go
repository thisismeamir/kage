package atoms

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	"github.com/thisismeamir/kage/pkg/form"
)

// DeleteNodePathResponse is the response structure for removing an node path.
type DeleteNodePathResponse struct {
	NodePath string `json:"node_path"`
	Deleted  bool   `json:"removed"`
	Message  string `json:"message,omitempty" jsonschema:"omitempty" jsonschema_extras:"description=Message about the removal status"`
}

func DeleteAtomPath(c *gin.Context) {
	var req form.FormPath
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	nodePaths := i.GetGlobalConfig().NodePaths
	newNodePaths := []form.FormPath{}
	deleted := false
	for _, atomPath := range nodePaths {
		if atomPath.Path != req.Path {
			newNodePaths = append(newNodePaths, atomPath)
		} else {
			deleted = true
		}
	}

	if deleted {
		cfg := i.GetGlobalConfig()
		cfg.NodePaths = newNodePaths
		i.SetGlobalConfig(cfg)
		i.SaveConfigFile()
		resp := DeleteNodePathResponse{
			NodePath: req.Path,
			Deleted:  true,
			Message:  "Node path deleted successfully.",
		}
		c.JSON(200, resp)
	} else {
		c.JSON(404, gin.H{"error": "node path not found"})
	}
}

func deleteModulePath(c *gin.Context) {
	var req form.FormPath
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	atomPaths := i.GetGlobalConfig().NodePaths
	newAtomPaths := []form.FormPath{}
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
		cfg.NodePaths = newAtomPaths
		i.SetGlobalConfig(cfg)
		i.SaveConfigFile()
		resp := DeleteNodePathResponse{
			NodePath: req.Path,
			Deleted:  true,
			Message:  "Node path deleted successfully.",
		}
		c.JSON(200, resp)
	} else {
		c.JSON(404, gin.H{"error": "node path not found"})
	}
}
