package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	RedisAddr  string `yaml:"redis_addr"`
	Db         string `yaml:"dsn"`
	ServerAddr string `yaml:"local_grpc_addr"`
}

func (c *Config) RedisAddress() string {
	return c.RedisAddr
}

func (c *Config) DSN() string {
	return c.Db
}

func (c *Config) ServiceAddr() string {
	return c.ServerAddr
}

func LoadConfig(fpath string) (*Config, error) {
	fpath, err := filepath.Abs(fpath)
	if err != nil {
		return nil, err
	}

	fd, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	res := &Config{}
	err = yaml.NewDecoder(fd).Decode(res)

	// log.Println(res.db, res.redisAddr, res.serverAddr)
	return res, err

}
