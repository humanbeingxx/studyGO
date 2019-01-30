package web

import (
	"github.com/kataras/iris"
)

func CreateApp() *iris.Application {
	app := iris.New()
	app.PartyFunc("/transform", func(child iris.Party) {
		child.Get("/", func(ctx iris.Context) {
			ctx.Write([]byte("Transform !!!"))
		})
		child.Get("/head", func(ctx iris.Context) {
			ctx.Write([]byte("I'm Head !!!"))
		})
		child.Get("/body", func(ctx iris.Context) {
			ctx.Write([]byte("I'm Body !!!"))
		})
		child.Get("/foot", func(ctx iris.Context) {
			ctx.Write([]byte("I'm Feet !!!"))
		})
	})
	return app
}
