package web

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
)

func CreateApp() *iris.Application {
	app := iris.New()

	app.RegisterView(iris.HTML("/Users/xiaoshuang.cui/work/github/studyGO/src/web/views", ".html"))

	app.Get("/index", func(ctx iris.Context) {
		ctx.ViewData("message", "Welcome to Seibertron")
		ctx.View("welcome.html")
	})

	transform(app)
	return app
}

func transform(app *iris.Application) *iris.Application {
	authConfig := basicauth.Config{
		Users:   map[string]string{"OptimusPrime": "Freedom"},
		Realm:   "Authorization Required",
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	needAuth := app.Party("/transform", authentication)

	needAuth.Get("/", func(ctx iris.Context) {
		ctx.Write([]byte("Transform !!!"))
	})
	needAuth.Get("/head", func(ctx iris.Context) {
		ctx.Write([]byte("I'm Head !!!"))
	})
	needAuth.Get("/body", func(ctx iris.Context) {
		ctx.Write([]byte("I'm Body !!!"))
	})
	needAuth.Get("/foot", func(ctx iris.Context) {
		ctx.Write([]byte("I'm Feet !!!"))
	})
	return app
}
