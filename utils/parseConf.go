package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type DbConf struct {
	Mysql struct {
		host     string `yaml:"host"`
		port     string `yaml:"port"`
		user     string `yaml:"user"`
		password string `yaml:"password"`
		db       string `yaml:"db"`
	}
	Redis struct {
		host string
		port int
		auth string
	}
}

type Conf struct {
	MysqlDsn string
	RedisDsn string
}

func (c *Conf) parseConf() {
	conf := new(DbConf)
	confFile, err := ioutil.ReadFile("./conf/database.yaml")

	if err != nil {
		log.Printf("无法读取配置文件 #%v ", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(confFile, conf)
	if err != nil {
		log.Printf("解析配置文件出错 #%v ", err)
		os.Exit(1)
	}

	fmt.Println(conf)
}
