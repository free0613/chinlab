package base

import (
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"resk.micro/infra"
	"resk.micro/tracer"
)

var irisapp *iris.Application

type IrisStarter struct {
	infra.BaseStarter
}

func Iris() *iris.Application {
	return irisapp
}

func (app *IrisStarter) Init(ctx infra.StarterContext) {
	irisapp = irisInit()

}

func (app *IrisStarter) Setup(ctx infra.StarterContext) {
	conf := GetProps()
	port, _ := conf.String("appinfo.port")
	irisapp.Run(iris.Addr(":"+port), iris.WithCharset("UTF-8"))
}

func (app *IrisStarter) StartBlocking() bool {
	return true
}
func irisInit() *iris.Application {
	app := iris.New()
	ApiInit(app)
	app.Use(tracer.Tracing())

	return app
}

func makeAccessLog() *accesslog.AccessLog {
	log := accesslog.File("access.log")
	log.AddOutput(os.Stdout)

	log.Delim = '|'
	log.TimeFormat = "2020-01-01 13:23:31 223"
	log.IP = true
	log.BytesReceivedBody = true
	log.BytesSentBody = true
	log.BytesSent = false
	log.BytesReceived = true
	log.RequestBody = true
	log.ResponseBody = false
	log.PanicLog = accesslog.LogHandler

	log.SetFormatter(&accesslog.JSON{
		Indent:    " ",
		HumanTime: true,
	})

	return log
}
