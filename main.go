package main

import (
	_ "go-raspberrypi/boot"
	_ "go-raspberrypi/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
