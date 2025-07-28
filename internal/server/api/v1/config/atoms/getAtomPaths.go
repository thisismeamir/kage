package atoms

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	"github.com/thisismeamir/kage/pkg/form"
)

func GetAtomPaths(c *gin.Context) {
	resp := i.GetGlobalConfig().NodePaths
	if resp == nil {
		resp = []form.FormPath{}
		c.JSON(200, resp)
	}
	c.JSON(200, resp)
}
