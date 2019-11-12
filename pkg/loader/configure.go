package configure

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type ConfLoader interface {
	loadOptions()
}

/*
* ConfFile this struct and app.yml corresponding.
* Mysql database link options.
* Redis database link options.
* Http http server options.
* Jwt authentication params.
 */
type ConfFile struct {
	App struct {
		RunMode string `yaml:"run_mode"`
		DbType  string `yaml:"database_type"`
	}
	Mysql struct {
		Host          string `yaml:"host"`
		Port          int    `yaml:"port"`
		User          string `yaml:"user"`
		Password      string `yaml:"password"`
		Charset       string `yaml:"charset"`
		ParseTime     bool   `yaml:"parseTime"`
		Location      string `yaml:"loc"`
		Db            string `yaml:"db"`
		DbPre         string `yaml:"db_pre"`
		MaxIdleConn   int    `yaml:"max_idle_conn"`
		MaxOpenConn   int    `yaml:"max_open_conn"`
		SingularTable bool   `yaml:"singular_table_name"`
	}
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Auth     string `yaml:"auth"`
		Protocol string `yaml:"protocol"`
		Db       int    `yaml:"db"`
	}
	Server struct {
		Host            string        `yaml:"host"`
		Port            int           `yaml:"port"`
		ReadTimeout     time.Duration `yaml:"read_timeout"`
		WriteTimeout    time.Duration `yaml:"write_timeout"`
		IdleTimeout     time.Duration `yaml:"idle_timeout"`
		MaxHeaderBytes  int           `yaml:"max_header_bytes"`
		EnableTLS       bool          `yaml:"enable_tls"`
		SSLCertfilePath string        `yaml:"ssl_certfile_path"`
		SSLKeyfilePath  string        `yaml:"ssl_keyfile_path"`
	}
	Jwt struct {
		Secret string `yaml:"secret"`
	}
}

//conf type of *ConfFile, this variable storage app.yml contents.
var conf *ConfFile

//app options
type AppOptions struct {
	RunMode string
	DbType  string
}

//this method implements ConfLoader interface.
//load app options this method load the Gin framework runtime options.
func (c *AppOptions) loadOptions() {
	c.RunMode = conf.App.RunMode
	c.DbType = conf.App.DbType
}

//mysql options
type MysqlOptions struct {
	Dsn           string
	DbPre         string
	MaxIdleConn   int
	MaxOpenConn   int
	SingularTable bool
}

//this method implements ConfLoader interface.
//load mysql options this method load mysql runtime options.
func (c *MysqlOptions) loadOptions() {
	c.Dsn = fmt.Sprintf("%s:%s@(%s:%d)/%s%s?charset=%s&parseTime=%t&loc=%s", conf.Mysql.User, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.DbPre, conf.Mysql.Db, conf.Mysql.Charset, conf.Mysql.ParseTime, conf.Mysql.Location)
	c.DbPre = conf.Mysql.DbPre
	c.MaxIdleConn = conf.Mysql.MaxIdleConn
	c.MaxOpenConn = conf.Mysql.MaxOpenConn
	c.SingularTable = conf.Mysql.SingularTable
}

//redis options
type RedisOptions struct {
	Addr     string
	Password string
	Db       int
}

//this method implements ConfLoader interface.
//load redis options this method load redis runtime options.
func (c *RedisOptions) loadOptions() {
	c.Addr = fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port)
	c.Password = conf.Redis.Auth
	c.Db = conf.Redis.Db
}

//server options
type ServerOptions struct {
	Addr            string
	EnableTLS       bool
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	MaxHeaderBytes  int
	SSLCertfilePath string
	SSLKeyfilePath  string
}

//this method implements ConfLoader interface.
//load server options this method load http server runtime option.
func (c *ServerOptions) loadOptions() {
	//adress and port
	c.Addr = fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	//read time out, unit(second)
	c.ReadTimeout = conf.Server.ReadTimeout * time.Second
	//write time out, unit(second)
	c.WriteTimeout = conf.Server.WriteTimeout * time.Second
	//idle time out, unit(second)
	c.IdleTimeout = conf.Server.IdleTimeout * time.Second
	//idle time out
	c.MaxHeaderBytes = conf.Server.MaxHeaderBytes
	//enable ssl
	c.EnableTLS = conf.Server.EnableTLS
	//ssl cert file path
	c.SSLCertfilePath = conf.Server.SSLCertfilePath
	//ssl key file path
	c.SSLKeyfilePath = conf.Server.SSLKeyfilePath
}

//jwt options
type JwtOptions struct {
	Secret string
}

//this method implements ConfLoader interface.
//load jwt Options.
func (c *JwtOptions) loadOptions() {
	c.Secret = conf.Jwt.Secret
}

var (
	AppConf    AppOptions
	MysqlConf  MysqlOptions
	RedisConf  RedisOptions
	ServerConf ServerOptions
	JwtConf    JwtOptions
)

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

	//combination struct and check type asserts
	AllOptionsAsserts := []struct {
		ConfLoader
	}{
		{&AppConf},
		{&MysqlConf},
		{&RedisConf},
		{&ServerConf},
		{&JwtConf},
	}

	//load all options
	for _, o := range AllOptionsAsserts {
		o.loadOptions()
	}
}
