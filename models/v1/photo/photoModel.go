package photo

import (
	configure "a6-api/utils/loader"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Photo struct{}

func (p *Photo) TablePrefix() string {
	path := reflect.TypeOf(Photo{}).PkgPath()
	packagePathArray := strings.Split(path, "/")
	prefix := packagePathArray[len(packagePathArray)-1]
	return prefix
}

func init() {
	fmt.Println(configure.AppConf.DbType)
	db, err := gorm.Open(configure.AppConf.DbType, configure.MysqlConf.Dsn[1])

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return configure.MysqlConf.DbPre + defaultTableName
	}

	//table name not use plural
	db.SingularTable(configure.MysqlConf.SingularTable)

	/* db link pool BEGIN */
	//max idle connection numbers
	db.DB().SetMaxIdleConns(configure.MysqlConf.MaxIdleConn)
	//max open connection numbers
	db.DB().SetMaxOpenConns(configure.MysqlConf.MaxOpenConn)
	/* db link pool END */
}

func CloseDB() {
	defer db.Close()
}
