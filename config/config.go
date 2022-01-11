package config

import (
	"os"
	"strings"

	"github.com/0w0mewo/budong/internal/persistent"
)

type Config struct {
	redisAddr   string
	db          string
	serverAddr  string
	storageType string
}

func (c *Config) RedisAddress() string {
	return c.redisAddr
}

func (c *Config) DSN() string {
	return c.db
}

func (c *Config) ServiceAddr() string {
	return c.serverAddr
}

func (c *Config) StoreType() persistent.RepoProvider {
	switch t := strings.ToLower(c.storageType); t {
	case "sqlite":
		return persistent.SQLITE
	case "mongo":
		return persistent.MONGO
	default:
		panic("unknown storage type")
	}
}

func LoadConfig() *Config {

	res := &Config{}

	res.redisAddr = os.Getenv("REDIS_ADDR")
	res.db = os.Getenv("DB_DSN")
	res.serverAddr = os.Getenv("LOCAL_ADDR")
	res.storageType = os.Getenv("DB_TYPE")

	if res.redisAddr == "" {
		res.redisAddr = "localhost"
	}

	if res.serverAddr == "" {
		res.serverAddr = ":9999"
	}

	if res.storageType == "" || res.db == "" {
		panic("unknown DB config")
	}

	return res

}
