package config

import (
	"log"
	"strconv"

	"github.com/Zenk41/sipencari-rest-api/util"
	"golang.org/x/crypto/bcrypt"
)

type JWTConfig struct {
	JWT_SECRET_KEY   string
	JWT_COST         int
	JWT_EXP_DURATION int
}

func LoadJWTConfig() JWTConfig {
	cost, err := strconv.Atoi(util.GetEnv("JWT_COST"))
	if err != nil {
		log.Print("Invalid cost")
		cost = bcrypt.DefaultCost
	}
	exp, err := strconv.Atoi(util.GetEnv("JWT_EXP_DURATION"))
	if err != nil {
		log.Print("Invalid duration")
		cost = 1
	}
	return JWTConfig{
		JWT_SECRET_KEY:   util.GetEnv("JWT_SECRET_KEY"),
		JWT_COST:         cost,
		JWT_EXP_DURATION: exp,
	}
}
