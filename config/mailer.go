package config

import "github.com/Zenk41/sipencari-rest-api/util"

type mailConfig struct {
	USERNAME string
	PASSWORD string
	HOST     string
	PORT     string
	TIMEOUT  string
	SENDER   string
}

func LoadMailConfig() mailConfig {
	return mailConfig{
		USERNAME: util.GetEnv("USERNAME"),
		PASSWORD: util.GetEnv("PASSWORD"),
		HOST:     util.GetEnv("HOST"),
		PORT:     util.GetEnv("PORT"),
		TIMEOUT:  util.GetEnv("TIMEOUT"),
		SENDER:   util.GetEnv("NAME"),
	}
}
