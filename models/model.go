package model

import (
	configure "a6-api/pkg/loader"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Modeler interface {
	List() string
	Detail() string
}

type Model struct{}

func (model *Model) List(m Modeler) string {
	return m.List()
}

func (model *Model) Detail(m Modeler) string {
	return m.Detail()
}

func init() {
	db, err := gorm.Open(configure.AppConf.DbType, configure.MysqlConf.Dsn)

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return configure.MysqlConf.DbPre + defaultTableName
	}

	//table name not use plural
	db.SingularTable(configure.MysqlConf.SingularTable)
	//max idle connection numbers
	db.DB().SetMaxIdleConns(configure.MysqlConf.MaxIdleConn)
	//max open connection numbers
	db.DB().SetMaxOpenConns(configure.MysqlConf.MaxOpenConn)
}

func CloseDB() {
	defer db.Close()
}
