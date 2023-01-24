package config

import ()

type LoggerConfig struct {
	Format string
}

func LoadLoggerConfig() LoggerConfig {
	return LoggerConfig{
		Format: "[${time_rfc3339}] method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
	}
}
