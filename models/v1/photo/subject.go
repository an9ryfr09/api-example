package photo

import (
	"fmt"
	"reflect"
	"strings"
)

type Subject struct {
	id      int    `gorm:"id"`
	subject string `gorm:"subject"`
}

var (
	photo   *Photo
	subject *Subject
)

func (s *Subject) TableName() string {
	prefix := photo.TablePrefix()
	table := strings.ToLower(reflect.TypeOf(Subject{}).Name())
	return fmt.Sprintf("%s_%s", prefix, table)
}

func (s *Subject) List() string {
	return s.TableName()
}

func (s *Subject) Detail() string {
	return "subject list"
}
