package main

import (
	"github.com/kataras/iris"

	"ChineseChess/server/routers"
)

func main() {

	app := iris.New()

	routers.RouteV1(app)

	app.Run(iris.Addr(":6666"))
}
