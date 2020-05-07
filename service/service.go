package service

import (
	"../handler/user"
	"github.com/gin-gonic/gin"
)

func Run(addr string) {
	// gin
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/hello", user.HelloHandler)
		v1.GET("/users", user.UsersHandler)
	}

	router.Run(addr)
}
