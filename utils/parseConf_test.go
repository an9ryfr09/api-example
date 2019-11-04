package utils

import (
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestParseConf(t *testing.T) {
	t.Run("parse yaml file", func(t *testing.T) {
		dbconf := new(DbConf)
		//dbconf.parseConf
		confFile, _ := ioutil.ReadFile("./conf/database.yaml")
		yaml.Unmarshal(confFile, dbconf)
		t.Fatalf("dbconf %v", dbconf)
	})
}
