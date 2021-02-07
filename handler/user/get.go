package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func UsersHandler(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}
