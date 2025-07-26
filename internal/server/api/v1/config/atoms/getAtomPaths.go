package atoms

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	"github.com/thisismeamir/kage/pkg/atom"
)

func GetAtomPaths(c *gin.Context) {
	resp := i.GetGlobalConfig().AtomPaths
	if resp == nil {
		resp = []atom.AtomPath{}
		c.JSON(200, resp)
	}
	c.JSON(200, resp)
}
