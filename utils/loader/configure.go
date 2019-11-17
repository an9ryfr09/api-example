package configure

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type confLoader interface {
	loadOptions()
}

/*
* confFile this struct and app.yml corresponding.
* Core app runtime options
* Mysql database link options.
* Redis database link options.
* Http http server options.
* Jwt authentication params.
 */
type confFile struct {
	Core struct {
		RunMode    string `yaml:"run_mode"`
		DbType     string `yaml:"database_type"`
		CpuCoreNum uint8  `yaml:"cpu_core_num"`
		PerPageNum uint16 `yaml:"per_page_num"`
	}
	Mysql struct {
		Host          string        `yaml:"host"`
		Port          uint16        `yaml:"port"`
		User          string        `yaml:"user"`
		Password      string        `yaml:"password"`
		Charset       string        `yaml:"charset"`
		ParseTime     bool          `yaml:"parseTime"`
		Location      string        `yaml:"location"`
		Timeout       time.Duration `yaml:"timeout"`
		ReadTimeOut   time.Duration `yaml:"read_time_out"`
		WriteTimeOut  time.Duration `yaml:"write_time_out"`
		Db            []string      `yaml:"db"`
		DbPre         string        `yaml:"db_pre"`
		MaxIdleConn   uint16        `yaml:"max_idle_conn"`
		MaxOpenConn   uint16        `yaml:"max_open_conn"`
		SingularTable bool          `yaml:"singular_table_name"`
	}
	Redis struct {
		Host     string `yaml:"host"`
		Port     uint16 `yaml:"port"`
		Auth     string `yaml:"auth"`
		Protocol string `yaml:"protocol"`
		Db       uint8  `yaml:"db"`
	}
	Server struct {
		Host            string        `yaml:"host"`
		Port            uint16        `yaml:"port"`
		ReadTimeout     time.Duration `yaml:"read_timeout"`
		WriteTimeout    time.Duration `yaml:"write_timeout"`
		IdleTimeout     time.Duration `yaml:"idle_timeout"`
		MaxHeaderBytes  uint64        `yaml:"max_header_bytes"`
		EnableTLS       bool          `yaml:"enable_tls"`
		SSLCertfilePath string        `yaml:"ssl_certfile_path"`
		SSLKeyfilePath  string        `yaml:"ssl_keyfile_path"`
	}
	Jwt struct {
		Secret string `yaml:"secret"`
	}
	Log struct {
		Path      string        `yaml:"path"`
		MaxAge    time.Duration `yaml:"max_age"`
		SplitTime time.Duration `yaml:"split_time"`
	}
}

//conf type of *ConfFile, this variable storage app.yml contents.
var conf *confFile

//core options
type coreOptions struct {
	RunMode    string
	DbType     string
	CpuCoreNum uint8
	PerPageNum uint16
	LogPath    string
}

//this method implements confLoader interface.
//load app options this method load the Gin framework runtime options.
func (c *coreOptions) loadOptions() {
	c.CpuCoreNum = conf.Core.CpuCoreNum
	c.RunMode = conf.Core.RunMode
	c.DbType = conf.Core.DbType
	c.PerPageNum = conf.Core.PerPageNum
}

//mysql options
type mysqlOptions struct {
	Dsn           []string
	Timeout       time.Duration
	ReadTimeOut   time.Duration
	WriteTimeOut  time.Duration
	DbPre         string
	MaxIdleConn   uint16
	MaxOpenConn   uint16
	SingularTable bool
}

//this method implements confLoader interface.
//load mysql options this method load mysql runtime options.
func (c *mysqlOptions) loadOptions() {
	for _, db := range conf.Mysql.Db {
		c.Dsn = append(c.Dsn, fmt.Sprintf("%s:%s@(%s:%d)/%s%s?charset=%s&parseTime=%t&loc=%s", conf.Mysql.User, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.DbPre, db, conf.Mysql.Charset, conf.Mysql.ParseTime, conf.Mysql.Location))
	}

	c.Timeout = conf.Mysql.Timeout * time.Second
	c.ReadTimeOut = conf.Mysql.ReadTimeOut * time.Second
	c.WriteTimeOut = conf.Mysql.WriteTimeOut * time.Second
	c.DbPre = conf.Mysql.DbPre
	c.MaxIdleConn = conf.Mysql.MaxIdleConn
	c.MaxOpenConn = conf.Mysql.MaxOpenConn
	c.SingularTable = conf.Mysql.SingularTable
}

//redis options
type redisOptions struct {
	Addr     string
	Password string
	Db       uint8
}

//this method implements confLoader interface.
//load redis options this method load redis runtime options.
func (c *redisOptions) loadOptions() {
	c.Addr = fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port)
	c.Password = conf.Redis.Auth
	c.Db = conf.Redis.Db
}

//server options
type serverOptions struct {
	Addr            string
	EnableTLS       bool
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	MaxHeaderBytes  uint64
	SSLCertfilePath string
	SSLKeyfilePath  string
}

//this method implements confLoader interface.
//load server options this method load http server runtime option.
func (c *serverOptions) loadOptions() {
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
type jwtOptions struct {
	Secret string
}

//this method implements confLoader interface.
//load jwt Options.
func (c *jwtOptions) loadOptions() {
	c.Secret = conf.Jwt.Secret
}

type logOptions struct {
	Path      string
	MaxAge    time.Duration
	SplitTime time.Duration
}

func (c *logOptions) loadOptions() {
	c.Path = conf.Log.Path
	c.MaxAge = conf.Log.MaxAge * time.Second * 86400
	c.SplitTime = conf.Log.SplitTime * time.Second * 86400
}

var (
	CoreConf   coreOptions
	MysqlConf  mysqlOptions
	RedisConf  redisOptions
	ServerConf serverOptions
	JwtConf    jwtOptions
	LogConf    logOptions
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
		confLoader
	}{
		{&CoreConf},
		{&MysqlConf},
		{&RedisConf},
		{&ServerConf},
		{&JwtConf},
		{&LogConf},
	}

	//load all options
	for _, loader := range AllOptionsAsserts {
		loader.loadOptions()
	}
}
