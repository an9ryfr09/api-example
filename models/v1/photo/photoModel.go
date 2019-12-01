package photo

import (
	"a6-api/utils/loader"
	"fmt"
	"log"
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
	photo := Photo{}
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s%s?charset=%s&parseTime=%t&loc=%s", loader.Load().Mysql.User, loader.Load().Mysql.Password, loader.Load().Mysql.Host, loader.Load().Mysql.Port, loader.Load().Mysql.DbPre, photo.TablePrefix(), loader.Load().Mysql.Charset, loader.Load().Mysql.ParseTime, loader.Load().Mysql.Location)
	db, err = gorm.Open(loader.Load().Core.DbType, dsn)

	if err != nil {
		log.Fatal(err.Error())
	}

	//enabled db log
	db.LogMode(true)

	//table name not use plural
	db.SingularTable(loader.Load().Mysql.SingularTable)

	/* db link pool BEGIN */
	//max idle connection numbers
	db.DB().SetMaxIdleConns(int(loader.Load().Mysql.MaxIdleConn))
	//max open connection numbers
	db.DB().SetMaxOpenConns(int(loader.Load().Mysql.MaxOpenConn))
	/* db link pool END */
}

func CloseDB() {
	defer db.Close()
}
