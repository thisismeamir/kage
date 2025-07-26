package modules

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/init"
	"github.com/thisismeamir/kage/internal/models"
)

func GetModulePaths(c *gin.Context) {
	resp := i.GetGlobalConfig().ModulePaths
	if resp == nil {
		resp = []models.ModulePath{}
		c.JSON(200, resp)
	}
	c.JSON(200, resp)
}
