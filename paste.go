package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"log"
)

var cache = make(map[string]*Paste)

type Paste struct {
	gorm.Model
	PasteId string `gorm:"PRIMARY_KEY"`
	Lang string
	Text string `gorm:"type:longtext"`
	Views int
	Likes int
}

func main() {
	db, err := gorm.Open("mysql", "root@/paste?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&Paste{})

	var pastes []*Paste
	db.Find(&pastes)

	for _, paste := range pastes {
		cache[paste.PasteId] = paste
	}

	fmt.Println("Loaded", len(cache), " pastes.")

	app := iris.New()
	engine := iris.Django("./pages", ".html")
	engine.Reload(true)
	app.RegisterView(engine)

	app.HandleDir("/public/js", "./pages/js/")

	HandleApis(app)

	app.Handle("GET", "/", func(ctx iris.Context) {
		err = ctx.View("index.html")
		if err != nil {
			log.Fatal(err)
		}
	})

	app.Handle("GET", "/{id:string}", func(ctx iris.Context) {
		id := ctx.Params().GetStringDefault("id", "")
		if id != "" {
			paste := GetPaste(id, ctx)

			paste.Views++;
			db.Save(&paste)

			ctx.ViewData("views", paste.Views)
			ctx.ViewData("likes", paste.Likes)
			ctx.ViewData("text", paste.Text)
			ctx.ViewData("lang", paste.Lang)
			ctx.ViewData("id", id)
			err = ctx.View("result.html")
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	})

	app.Handle("POST", "/create", func(ctx iris.Context) {
		id := ctx.Request().FormValue("id")
		custom := true
		if id == "" {
			id = String(16)
			custom = false
		}

		paste := Paste{
			PasteId: id,
			Text: ctx.Request().FormValue("text"),
			Lang: ctx.Request().FormValue("lang"),
			Views: 0,
			Likes: 0,
		}
		if cache[id] != nil && custom {
			ctx.ViewData("error", "Paste with this name already exists.")
			err = ctx.View("error.html")
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		db.Create(&paste)
		cache[paste.PasteId] = &paste
		ctx.Redirect("/" + id)
	})

	app.Handle("POST", "/like", func(ctx iris.Context) {
		id := ctx.Request().FormValue("id")
		paste := GetPaste(id, ctx)
		paste.Likes++;
		ctx.Redirect("/" + id)
	})

	err = app.Run(iris.Addr(":8080"))
	if err != nil {
		log.Fatal(err)
	}
}