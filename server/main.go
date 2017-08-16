package main

import (
	"github.com/kataras/iris"

	"ChineseChess/server/routers"
)

func main() {

	app := iris.New()

	routers.RouteV1(app)

	if err := app.Run(iris.Addr(":6666"), iris.WithCharset("UTF-8")); err != nil {
		panic(err)
	}
}
