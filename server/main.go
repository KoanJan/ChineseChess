package main

import (
	"ChineseChess/server/routers"
)

func main() {

	routers.StartTCP(6666)
}
