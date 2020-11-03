package config

import (
    "gopkg.in/ini.v1"
)

type Config struct{
	Port string
}

func Load_conf() Config {
    c, _ := ini.Load("conf.ini")
    conf := Config{
		Port: c.Section("server").Key("port").String(),
	}
	return conf
}
