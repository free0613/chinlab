package infra

import "github.com/gookit/ini"

const (
	KeyProps = "_conf"
)

type StarterContext map[string]interface{}

func (s StarterContext) Props() *ini.Ini {
	conf := s[KeyProps]
	if conf == nil {
		panic("config is not init")
	}
	return conf.(*ini.Ini)
}

func (s StarterContext) SetProps(conf *ini.Ini) {
	s[KeyProps] = conf
}

type Starter interface {
	Init(StarterContext)

	Setup(StarterContext)

	Start(StarterContext)

	StartBlocking() bool

	Stop(StarterContext)
}

type BaseStarter struct{}

func (b BaseStarter) Init(StarterContext) {}

func (b *BaseStarter) Setup(StarterContext) {}

func (b *BaseStarter) Start(StarterContext) {}

func (b *BaseStarter) StartBlocking() bool {
	return false
}

func (b *BaseStarter) Stop(StarterContext) {}

type startRegister struct {
	starters []Starter
}

var RegStart *startRegister = new(startRegister)

func (r *startRegister) Register(s Starter) {
	r.starters = append(r.starters, s)
}

func (r *startRegister) AllRegister() []Starter {
	return r.starters
}

func Register(s Starter) {
	RegStart.Register(s)
}
