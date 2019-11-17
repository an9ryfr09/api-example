package photo

import (
	configure "a6-api/utils/loader"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	var err error
	db, err = gorm.Open(configure.CoreConf.DbType, configure.MysqlConf.Dsn[1])

	if err != nil {

	}

	//enabled db log
	db.LogMode(true)

	//table name not use plural
	db.SingularTable(configure.MysqlConf.SingularTable)

	/* db link pool BEGIN */
	//max idle connection numbers
	db.DB().SetMaxIdleConns(int(configure.MysqlConf.MaxIdleConn))
	//max open connection numbers
	db.DB().SetMaxOpenConns(int(configure.MysqlConf.MaxOpenConn))
	/* db link pool END */
}

func CloseDB() {
	defer db.Close()
}
