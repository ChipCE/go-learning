package main

import "gin-example/plugins/plugina"
import "gin-example/plugins/pluginb"

func main() {
	go plugina.StartOne()
	pluginb.StartTwo()
}
