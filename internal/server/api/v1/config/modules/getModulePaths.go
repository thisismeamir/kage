package modules

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	. "github.com/thisismeamir/kage/pkg/graph"
)

func GetModulePaths(c *gin.Context) {
	resp := i.GetGlobalConfig().GraphPaths
	if resp == nil {
		resp = []GraphPath{}
		c.JSON(200, resp)
	}
	c.JSON(200, resp)
}
