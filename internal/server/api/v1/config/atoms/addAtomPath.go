package atoms

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	"github.com/thisismeamir/kage/pkg/atom"
	"os"
)

// AddAtomPathResponse is the response structure for checking the existence and validity of an atom path.
type AddAtomPathResponse struct {
	AtomPath string `json:"atomPath"`
	Added    bool   `json:"added"`
	Message  string `json:"message"`
}

func AddAtomPath(c *gin.Context) {
	var req atom.AtomPath

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
		for _, atomPath := range i.GetGlobalConfig().AtomPaths {
			if atomPath.Path == req.Path {
				c.JSON(200, gin.H{"error": "atom path already exists"})
				return
			}
		}
		// Add the new atom path to the global config
		i.SetGlobalConfig(i.Config{
			Name:        i.GetGlobalConfig().Name,
			BasePath:    i.GetGlobalConfig().BasePath,
			ModulePaths: i.GetGlobalConfig().ModulePaths,
			AtomPaths:   append(i.GetGlobalConfig().AtomPaths, req),
			Version:     i.GetGlobalConfig().Version,
			Server:      i.GetGlobalConfig().Server,
			Client:      i.GetGlobalConfig().Client,
		})

		i.SaveConfigFile()

	}
	// Example response (replace with your logic)
	resp := AddAtomPathResponse{
		AtomPath: req.Path,
		Added:    true,
		Message:  "Atom path added successfully.",
	}
	c.JSON(200, resp)
}
