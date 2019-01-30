package main

import (
	"web"

	"github.com/kataras/iris"
)

func main() {
	app := web.CreateApp()

	app.Run(iris.Addr("localhost:8081"))
}
