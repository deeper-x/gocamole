package main

import (
	"log"
	"os"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./tmpl", ".html").Reload(true))
	app.StaticWeb("/static", "./static")

	app.Handle("GET", "/getcontextstatus/{id:uint64}", func(ctx iris.Context) {
		ctx.HTML("#TODO")
	})

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.View("main.html")
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}

func init() {
	gopath := os.Getenv("GOPATH")

	if gopath == "" {
		log.Fatal("${GOPATH} not set. Please fix it.")
	}
}
