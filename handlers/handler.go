package handler

import (
	configure "a6-api/packages/conf"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Moder interface {
	List()
	Detail()
}

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}

func init() {
	db, err := gorm.Open(configure.AppConf.DbType, configure.MysqlConf.Dsn)

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return configure.MysqlConf.TablePre + defaultTableName
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
