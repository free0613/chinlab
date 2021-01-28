package infra

import (
	"github.com/gookit/ini"
)

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

func (b *BootApplication) init() {
	for _, start := range RegStart.AllRegister() {
		start.Init(b.startctx)
	}
}

func (b *BootApplication) setup() {
	for _, start := range RegStart.AllRegister() {
		start.Setup(b.startctx)
	}

}

func (b *BootApplication) start() {
	for i, start := range RegStart.AllRegister() {
		if start.StartBlocking() {
			if i+1 == len(RegStart.AllRegister()) {
				start.Start(b.startctx)
			} else {
				go start.Start(b.startctx)
			}
		} else {
			start.Start(b.startctx)
		}
	}
}

func (b *BootApplication) Startup() {
	b.init()
	b.start()
	b.setup()
}
