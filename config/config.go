package config

var GlobalConfig *config

func init() {
	GlobalConfig = new(config)
	GlobalConfig.redisAddr = "127.0.0.1:6379"
	GlobalConfig.db = "mongodb://127.0.0.1:27017"
	GlobalConfig.serverAddr = "127.0.0.1:9999"
}

type config struct {
	redisAddr  string
	db         string
	serverAddr string
}

func (c *config) RedisAddr() string {
	return c.redisAddr
}

func (c *config) DB() string {
	return c.db
}

func (c *config) Addr() string {
	return c.serverAddr
}
