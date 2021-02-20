package resk

import (
	"fmt"

	"resk.micro/infra"
	"resk.micro/infra/base"
)

func init() {
	fmt.Println("注册")
	infra.Register(&base.EtcdStarter{})
	infra.Register(&base.LogStarter{})
	infra.Register(&base.PropStarter{})
	infra.Register(&base.MysqlStarter{})
	infra.Register(&base.IrisStarter{})
}
