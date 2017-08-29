package main

import (
	"github.com/gin-gonic/gin"

	"ChineseChess/server/logger"
	"ChineseChess/server/routers"
)

func main() {

	engine := gin.New()
	engine.Use(gin.LoggerWithWriter(logger.LogFileWriter), gin.RecoveryWithWriter(logger.LogFileWriter))
	defer logger.LogFileWriter.Close()

	routers.Route(engine)

	engine.Run(":9088")
}
