package conf

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

//ConfFile this struct and app.yml corresponding.
//Mysql database link options
//Redis database link options
//Http http server options
//Jwt authentication params
type ConfFile struct {
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
	Http struct {
		Host           string `yaml:"host"`
		Port           int    `yaml:"port"`
		ReadTimeOut    int    `yaml:"read_timeout"`
		WriteTimeOut   int    `yaml:"write_timeout"`
		MaxHeaderBytes uint64 `yaml:"max_header_bytes"`
	}
	Jwt struct {
		Secret string `yaml:"secret"`
	}
}

//AppConf all config options in this struct.
type AppConf struct {
	MysqlDsn     string
	RedisOptions struct {
		Addr     string
		Password string
		Db       int
	}
	HttpOptions struct {
		Addr           string
		ReadTimeout    int
		WriteTimeout   int
		MaxHeaderBytes uint64
	}
	JwtOptions struct {
		Secret string
	}
}

//conf type is *ConfFile, this variable storage app.yml contents.
var conf *ConfFile

//init load app.yml and parsing config file.
func init() {
	confFile, err := ioutil.ReadFile("conf/app.yml")

	if err != nil {
		log.Printf("load file app.yml fail! #%v ", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(confFile, &conf)
	if err != nil {
		log.Printf("parsing file app.yml fail! #%v ", err)
		os.Exit(1)
	}
}

//loadMysqlDsn this method load mysql link string.
func (c *AppConf) loadMysqlDsn() string {
	return fmt.Sprintf("%s@(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s", conf.Mysql.User, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Db, conf.Mysql.Charset, conf.Mysql.ParseTime, conf.Mysql.Location)
}

//loadRedisOptions this method load redis config options。
func (c *AppConf) loadRedisOptions() {
	c.RedisOptions.Addr = fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port)
	c.RedisOptions.Password = conf.Redis.Auth
	c.RedisOptions.Db = conf.Redis.Db
}

//loadHttpOptions this method load http server options。
func (c *AppConf) loadHttpOptions() {
	c.HttpOptions.Addr = fmt.Sprintf("%s:%d", conf.Http.Host, conf.Http.Port)
	c.HttpOptions.ReadTimeout = conf.Http.ReadTimeOut
	c.HttpOptions.WriteTimeout = conf.Http.WriteTimeOut
	c.HttpOptions.MaxHeaderBytes = conf.Http.MaxHeaderBytes
}

//loadJwtOptions
func (c *AppConf) loadJwtOptions() {
	c.JwtOptions.Secret = conf.Jwt.Secret
}

//Load this method load all config。
func (c *AppConf) Load() {
	c.loadMysqlDsn()
	c.loadRedisOptions()
	c.loadHttpOptions()
	c.loadJwtOptions()
	fmt.Println(conf)
}
