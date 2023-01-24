package config

import "github.com/Zenk41/sipencari-rest-api/util"

type DBConfig struct {
	USERNAME string
	PASSWORD string
	NAME     string
	HOST     string
	PORT     string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		USERNAME: util.GetEnv("DB_USERNAME"),
		PASSWORD: util.GetEnv("DB_PASSWORD"),
		HOST:     util.GetEnv("DB_HOST"),
		PORT:     util.GetEnv("DB_PORT"),
		NAME:     util.GetEnv("DB_NAME"),
	}
}
