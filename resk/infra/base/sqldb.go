package base

import . "resk.micro/infra"

var databaseDbx *DBmsql

type DBmsql struct {
	BaseStarter
}

func (db DBmsql) Init(s StarterContext) {
	RegStart.Register(db)
}

func (db DBmsql) Start(s StarterContext) {

}
