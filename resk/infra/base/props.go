package base

import (
	"github.com/gookit/ini"
	"go.uber.org/zap"
	. "resk.micro/infra"
)

var props *ini.Ini

type PropStarter struct {
	BaseStarter
}

func GetProps() *ini.Ini {
	return props
}

func (p *PropStarter) Init(ctx StarterContext) {
	props = ctx.Props()
	logger.Info("初始化配置", zap.String("Name", "abc"))
}
