package config

import "fmt"

// RedisConfig contains the configuration for the Redis.
type RedisConfig struct {
	Host     string `default:"localhost"`
	Port     int16  `default:"6379"`
	Username string `default:""`
	Password string `default:""`
	Database int    `default:"0"`
}

// GetConnectionStringForDisplay returns a string to connect to the database with password redacted.
func (c *RedisConfig) GetConnectionStringForDisplay() string {
	return fmt.Sprintf("%s:%d (Db: %d)", c.Host, c.Port, c.Database)
}
