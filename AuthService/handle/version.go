package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func VersionEndPoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"version": "v0.0.0",
	})
}
