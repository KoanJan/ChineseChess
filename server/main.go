package main

import (
	"github.com/gin-gonic/gin"

	"ChineseChess/server/routers"
)

func main() {

	engine := gin.Default()

	routers.Route(engine)

	engine.Run(":9088")
}
