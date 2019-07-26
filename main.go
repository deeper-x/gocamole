package main

import (
	"fmt"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./tmpl", ".html").Reload(true))
	app.StaticWeb("/static", "./static")

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.ViewData("prova", "foo bar baz")
		ctx.View("main.html")
	})

	app.Handle("GET", "/get_state_data/{id:uint64}", func(ctx iris.Context) {
		idState := ctx.Params().Get("id")

		ctx.ViewData("id_passed", idState)
		ctx.View("main.html")
		fmt.Printf("id state: %v\n", idState)
	})

	app.Handle("GET", "/JSONCall/{id:uint64}", func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		ctx.JSON(iris.Map{"result": id})
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}
