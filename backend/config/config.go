package config

import (
	"log"

	"github.com/pelletier/go-toml"
)

type Config struct {
	listeningAddress string
	basepath         string
	dblink           string
}

var conf Config

func InitConfig(configPath string) {
	tom, err := toml.LoadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	conf.listeningAddress = tom.Get("listeningAddress").(string)
	conf.basepath = tom.Get("basepath").(string)
	conf.dblink = tom.Get("dblink").(string)
}

func ListeningAddress() string {
	return conf.listeningAddress
}

func BasePath() string {
	return conf.basepath
}

func DBLink() string {
	return conf.dblink
}
