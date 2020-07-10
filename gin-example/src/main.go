package main

import (
	"gin-example/plugins/plugina"
	"gin-example/plugins/pluginb"
)

func main() {
	go plugina.StartOne()
	pluginb.StartTwo()
}
