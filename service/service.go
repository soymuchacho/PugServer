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
		// 服务测试
		v1.GET("/version", user.HelloHandler) // 测试

		// 用户管理
		v1.GET("/users", user.GetUsersHandler)      // 获取所有的user
		v1.POST("/users", user.AddUserHandler)      // 新增一个user
		v1.PUT("/users", user.ModifyUserHandler)    // 修改一个user信息
		v1.DELETE("/users", user.DeleteUserHandler) // 删除一个user
	}

	router.Run(addr)
}
