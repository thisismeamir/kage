package modules

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	"github.com/thisismeamir/kage/pkg/module"
)

func GetModulePaths(c *gin.Context) {
	resp := i.GetGlobalConfig().ModulePaths
	if resp == nil {
		resp = []atom.ModulePath{}
		c.JSON(200, resp)
	}
	c.JSON(200, resp)
}
