package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	ServerConfig *ServerConfig
}

type ServerConfig struct {
	Port int
	Env  string
}

func GetConfig(env string) *Config {
	err := godotenv.Load(fmt.Sprintf("cfg_%s.env", env))
	if err != nil {
		panic(err)
	}

	return &Config{
		ServerConfig: &ServerConfig{
			Port: convertToInt(os.Getenv("SERVER_PORT"), 4000),
			Env:  env,
		},
	}
}

func convertToInt(str string, defaultVal int) int {
	if str == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultVal
	}
	return val
}
