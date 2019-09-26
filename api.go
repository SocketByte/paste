package main

import (
	"github.com/kataras/iris"
	"strconv"
)

func HandleApis(app *iris.Application) {
	app.Handle("GET", "/api/v1/get_views/{id:string}", func(ctx iris.Context) {
		id := ctx.Params().GetStringDefault("id", "")
		if id != "" {
			paste := GetPaste(id, ctx)
			_, _ = ctx.Text(strconv.Itoa(paste.Views))
		}
	})

	app.Handle("GET", "/api/v1/get_likes/{id:string}", func(ctx iris.Context) {
		id := ctx.Params().GetStringDefault("id", "")
		if id != "" {
			paste := GetPaste(id, ctx)
			_, _ = ctx.Text(strconv.Itoa(paste.Likes))
		}
	})
}
