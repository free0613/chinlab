package infra

import "github.com/gookit/ini"

type BootApplication struct {
	startctx StarterContext
	conf     *ini.Ini
}

func New(conf *ini.Ini) *BootApplication {
	b := &BootApplication{
		conf:     conf,
		startctx: StarterContext{},
	}
	b.startctx[KeyProps] = conf
	return b
}

func (b *BootApplication) Init() {
	for _, start := range RegStart.AllRegister() {
		start.Init(b.startctx)
	}
}

func (b *BootApplication) Setup() {
	panic("implement me")
}

func (b *BootApplication) Start() {
	panic("implement me")
}

func (b *BootApplication) Startup() {
	b.Init()
	b.Start()
	b.Setup()
}
