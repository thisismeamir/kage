package atoms

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/init"
	"github.com/thisismeamir/kage/internal/models"
)

func GetAtomPaths(c *gin.Context) {
	resp := i.GetGlobalConfig().AtomPaths
	if resp == nil {
		resp = []models.AtomPath{}
		c.JSON(200, resp)
	}
	c.JSON(200, resp)
}
