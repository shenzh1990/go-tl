package settings

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

var config = flag.String("f", "./pkg/settings/config.yaml", "config file")

type Config struct {
	Base  BaseConfig  `yaml:"Base"`
	Db    DbConfig    `yaml:"Db"`
	Redis RedisConfig `yaml:"Redis"`
}
type BaseConfig struct {
	RunMode       string        `yaml:"RunMode"`
	HTTPPort      int           `yaml:"HTTPPort"`
	ReadTimeout   time.Duration `yaml:"ReadTimeout"`
	WriteTimeout  time.Duration `yaml:"WriteTimeout"`
	PageSize      int           `yaml:"PageSize"`
	JwtSecret     string        `yaml:"JwtSecret"`
	CpuMaxProcess int           `yaml:"CpuMaxProcess"`
}
type DbConfig struct {
	Prefix     string `yaml:"Prefix"`
	DriverName string `yaml:"DriverName"`
	DBUrl      string `yaml:"DBUrl"`
}
type RedisConfig struct {
	RedisHost string `yaml:"RedisHost"`
	RedisDB   string `yaml:"RedisDB"`
	RedisPwd  string `yaml:"RedisPwd"`
	Timeout   int64  `yaml:"Timeout"`

	PoolMaxIdle     int   `yaml:"PoolMaxIdle"`
	PoolMaxActive   int   `yaml:"PoolMaxActive"`
	PoolIdleTimeout int64 `yaml:"PoolIdleTimeout"`
	PoolWait        bool  `yaml:"PoolWait"`
}

func (c *Config) getConf(filepath string) *Config {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
func init() {
	flag.Parse()
	CommonConfig.getConf(*config)
	LoadBase()
}

var (
	CommonConfig Config
	RunMode      string

	HTTPPort      int
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	CpuMaxProcess int
	PageSize      int
	JwtSecret     string
)

func LoadBase() {
	CpuMaxProcess = CommonConfig.Base.CpuMaxProcess
	RunMode = CommonConfig.Base.RunMode
	HTTPPort = CommonConfig.Base.HTTPPort
	ReadTimeout = CommonConfig.Base.ReadTimeout * time.Second
	WriteTimeout = CommonConfig.Base.WriteTimeout * time.Second
	JwtSecret = CommonConfig.Base.JwtSecret
	PageSize = CommonConfig.Base.PageSize
}
