package main

import (
	"github.com/gin-gonic/gin"

	"ChineseChess/server/logger"
	"ChineseChess/server/routers"
)

func main() {

	engine := gin.New()
	engine.Use(gin.LoggerWithWriter(logger.Writer), gin.RecoveryWithWriter(logger.Writer))
	defer logger.Writer.Close()

	routers.Route(engine)

	engine.Run(":9088")
}
