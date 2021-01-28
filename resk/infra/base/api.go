package base

import (
	"fmt"

	"github.com/kataras/iris/v12"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	. "resk.micro/internal/api"
)

func AuthRequired(ctx iris.Context) {
	fmt.Println("adfasf")
}

func ApiInit(app *iris.Application) {
	// log := makeAccessLog()
	// defer log.Close()
	app.UseRouter(recover2.New())

	authorized := app.Party("/")
	authorized.Use(AuthRequired)
	{
		authorized.Get("/login", GetUserInfo)

	}

}
