package base

import (
	"github.com/gookit/ini"
	. "resk.micro/infra"
)

var props *ini.Ini

type prop struct {
	BaseStarter
}

func Getprops() *ini.Ini {
	return props
}

func (p *prop) Init(ctx StarterContext) {
	props = ctx.Props()

}
