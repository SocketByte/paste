package main

import (
	"github.com/kataras/iris"
	"log"
)

func GetPaste(id string, ctx iris.Context) *Paste {
	paste := cache[id]
	if paste == nil {
		ctx.ViewData("error", "This paste does not exist.")
		err := ctx.View("error.html")
		if err != nil {
			log.Fatal(err)
		}
		_ = ctx.View("index.html")
		return nil
	}
	return paste
}
