package version

import (
	"PubServer/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VersionHandler(c *gin.Context) {
	version := "no version"
	// 获取版本号
	dir, err := GetCurrentDir()
	if err != nil {

	}
	c.String(http.StatusOK, "")
}
