package routers

import (
	"ChineseChess/server/routers/core"
)

func StartTCP(port int) {

	core.ServeTCP(port, Router())
}
