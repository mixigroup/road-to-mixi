package configs

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"sync"
)

var (
	conf Config
	once sync.Once
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port int `default:"1323"`
}

type DBConfig struct {
	Driver     string `default:"mysql"`
	DataSource string `default:"root:@(db:3306)/app"`
}

func Get() Config {
	once.Do(func() {
		if err := envconfig.Process("server", &conf.Server); err != nil {
			log.Fatal(err.Error())
		}
		if err := envconfig.Process("db", &conf.DB); err != nil {
			log.Fatal(err.Error())
		}
	})
	return conf
}
