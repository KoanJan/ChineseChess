package routers

import (
	"github.com/gin-gonic/gin"

	"ChineseChess/server/routers/middlewares"
	"ChineseChess/server/routers/v1"
)

func routeV1(e *gin.Engine) {

	// v1接口组
	group := e.Group("/api/v1")

	// 中间件
	group.Use(middlewares.Handlers...)

	// 路由表
	group.GET("/hello", v1.Hello)

	// user
	group.POST("/user", v1.CreateUser)
	group.POST("/session", v1.Login)
	group.DELETE("/sesssion", v1.Logout)

}
