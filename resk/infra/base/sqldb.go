package base

import (
	"context"
	"database/sql"
	. "resk.micro/infra"
)

var dbx *sql.DB
var _mapping = new(*sql.DB)

type Mysql struct {
	BaseStarter
}

func GetMysqlConn(ctx context.Context) *sql.DB {
	return dbx
}

func (m *Mysql) Setup(ctx StarterContext) {
	/*	conf :=ctx.Props()
		host, ok := conf.String("mysql.host")
		dbx*/
}

func Entermapping() {

}
