package config

import (
	"github.com/gin-gonic/gin"
	i "github.com/thisismeamir/kage/internal/bootstrap/config"
)

func GetConfiguration(c *gin.Context) {
	resp := i.GetGlobalConfig()
	c.JSON(200, resp)
}
