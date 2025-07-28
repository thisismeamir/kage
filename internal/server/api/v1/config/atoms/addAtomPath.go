package atoms

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	"github.com/thisismeamir/kage/pkg/form"
	"os"
)

// AddNodePathResponse is the response structure for checking the existence and validity of an node path.
type AddNodePathResponse struct {
	AtomPath string `json:"atomPath"`
	Added    bool   `json:"added"`
	Message  string `json:"message"`
}

func AddNodePath(c *gin.Context) {
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
		for _, atomPath := range i.GetGlobalConfig().NodePaths {
			if atomPath.Path == req.Path {
				c.JSON(200, gin.H{"error": "node path already exists"})
				return
			}
		}
		// Add the new node path to the global config
		i.SetGlobalConfig(i.Config{
			Name:       i.GetGlobalConfig().Name,
			BasePath:   i.GetGlobalConfig().BasePath,
			GraphPaths: i.GetGlobalConfig().GraphPaths,
			NodePaths:  append(i.GetGlobalConfig().NodePaths, req),
			Version:    i.GetGlobalConfig().Version,
			Server:     i.GetGlobalConfig().Server,
			Client:     i.GetGlobalConfig().Client,
		})

		i.SaveConfigFile()

	}
	// Example response (replace with your logic)
	resp := AddNodePathResponse{
		AtomPath: req.Path,
		Added:    true,
		Message:  "Atom path added successfully.",
	}
	c.JSON(200, resp)
}
