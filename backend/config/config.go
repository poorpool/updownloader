package config

import (
	"log"

	"github.com/pelletier/go-toml"
)

type Config struct {
	listeningAddress string
	webserverAddress string
	basepath         string
	dblink           string
	expireHourNum    int64
}

var conf Config

func InitConfig(configPath string) {
	tom, err := toml.LoadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	conf.listeningAddress = tom.Get("listeningAddress").(string)
	conf.webserverAddress = tom.Get("webserverAddress").(string)
	conf.basepath = tom.Get("basepath").(string)
	conf.dblink = tom.Get("dblink").(string)
	conf.expireHourNum = tom.Get("expireHourNum").(int64)
}

func ListeningAddress() string {
	return conf.listeningAddress
}

func WebServerAddress() string {
	return conf.webserverAddress
}

func BasePath() string {
	return conf.basepath
}

func DBLink() string {
	return conf.dblink
}

func ExpireHourNum() int64 {
	return conf.expireHourNum
}
