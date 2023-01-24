package util

import (
	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Msg("Error when read .env file")
	}
	return viper.GetString(key)
}
