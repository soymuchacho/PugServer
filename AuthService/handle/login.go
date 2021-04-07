package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginEndPoint(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("login test"))
}

func CheckAuthEndPoint(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("check auth"))
}
