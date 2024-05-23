package conf

import (
	"errors"

	"gopkg.in/ini.v1"
)

type Server struct {
	Host string
	Port int
}

type config struct {
	Server Server
	Openai string
}

var Config *config

var cachePath string

func Init() error {

	// Load main config.ini

	cfg, err := ini.Load("./config.ini")
	if err != nil {
		return err
	}

	// Parse server configuration

	var c config
	if c.Server.Host = cfg.Section("").Key("host").String(); len(c.Server.Host) == 0 {
		return errors.New("config error, field 'host' not found or not have values")
	}
	if c.Server.Port, err = cfg.Section("").Key("port").Int(); err != nil {
		return errors.New("config error, field 'port' not found or not have values")
	}
	Config = &c

	// Parse openai api

	if Config.Openai = cfg.Section("").Key("openai").String(); len(Config.Openai) == 0 {
		return errors.New("config error, field 'api' not found or not have values")
	}

	// Parse bots path

	if cachePath = cfg.Section("").Key("cache").String(); len(cachePath) == 0 {
		return errors.New("config error, field 'cache' not found or not have values")
	}

	return nil

}
