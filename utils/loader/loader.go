package loader

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

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
		Addr            string        `yaml:"addr"`
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
var (
	once sync.Once
	conf *confFile
	mu   sync.RWMutex
)

//Load parsing app.yml and load to struct in singleton
//add R-lock with non-conflict
func Load() *confFile {
	once.Do(Reload)
	mu.RLock()
	defer mu.RUnlock()
	return conf
}

//Reload re-parsing struct and load to struct in singleton
//add RW-lock with non-conflict
func Reload() {
	var config *confFile
	confFile, err := ioutil.ReadFile("conf/app.yml")
	if err != nil {
		log.Fatalf("load file app.yml fail! #%v ", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(confFile, &config)
	if err != nil {
		log.Fatalf("parsing file app.yml fail! #%v ", err)
		os.Exit(1)
	}

	mu.Lock()
	defer mu.Unlock()
	conf = config
}
