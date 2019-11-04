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
		Host      string   `yaml:"host"`
		Port      int      `yaml:"port"`
		User      string   `yaml:"user"`
		Password  string   `yaml:"password"`
		Charset   string   `yaml:"charset"`
		ParseTime bool     `yaml:"parseTime"`
		Location  string   `yaml:"loc"`
		Db        []string `yaml:"db"`
	}
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Auth     string `yaml:"auth"`
		Protocol string `yaml:"protocol"`
		Db       int    `yaml:"db"`
	}
}

type Conf struct {
	MysqlDsn     string
	RedisOptions struct {
		Addr     string
		Password string
		Db       int
	}
}

func (c *Conf) parseConfFile() *DbConf {
	var conf DbConf
	confFile, err := ioutil.ReadFile("./conf/database.yaml")

	if err != nil {
		log.Printf("无法读取配置文件 #%v ", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(confFile, &conf)
	if err != nil {
		log.Printf("解析配置文件出错 #%v ", err)
		os.Exit(1)
	}

	return &conf
}

func (c *Conf) LoadConf() {
	conf := c.parseConfFile()
	c.MysqlDsn = fmt.Sprintf("%s@(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s", conf.Mysql.User, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Db, conf.Mysql.Charset, conf.Mysql.ParseTime, conf.Mysql.Location)
	c.RedisOptions.Addr = fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port)
	c.RedisOptions.Password = conf.Redis.Auth
	c.RedisOptions.Db = conf.Redis.Db
}
