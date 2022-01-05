package main

import (
	"App-CloudBase-mcdull-mall/web"
)

func main() {
	app := &web.App{}
	app.Initialize()
	app.Run()
}
