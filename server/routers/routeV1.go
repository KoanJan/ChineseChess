package routers

import (
	"github.com/kataras/iris"

	"ChineseChess/server/routers/v1"
)

func RouteV1(app *iris.Application) {

	app.Get("/api/v1/hello", v1.Hello)

}
