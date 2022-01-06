package config

var GlobalConfig *config

func init() {
	GlobalConfig = new(config)
	GlobalConfig.redisAddr = "127.0.0.1:6379"
	GlobalConfig.db = "/tmp/test.db"
	GlobalConfig.httpAddr = "127.0.0.1:9999"
}

type config struct {
	redisAddr string
	db        string
	httpAddr  string
}

func (c *config) RedisAddr() string {
	return c.redisAddr
}

func (c *config) DB() string {
	return c.db
}

func (c *config) HttpAddr() string {
	return c.httpAddr
}
