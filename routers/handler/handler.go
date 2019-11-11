package handler

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Modeler interface {
	List() string
	Detail() string
}
