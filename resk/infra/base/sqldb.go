package base

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	. "resk.micro/infra"
)

var db *gorm.DB

type MysqlStarter struct {
	BaseStarter
}

func GetMysqlConn(ctx context.Context) *gorm.DB {
	return db
}

func (m *MysqlStarter) Setup(ctx StarterContext) {
	conf := ctx.Props()
	host, _ := conf.String("mysql.host")
	user, _ := conf.String("mysql.usernmae")
	pass, _ := conf.String("mysql.pass")
	dbname, _ := conf.String("mysql.dbname")

	dsnString := user + ":" + pass + "@tcp(" + host + ")/" + dbname + "?charset=utf8mb4&parseTime=true&loc=Local"

	dbx, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: dsnString,
	}), &gorm.Config{})

	sqlDb, err := dbx.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxIdleTime(60 * time.Second)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	logger.Info("connection is success,constring is")
	db = dbx

}
