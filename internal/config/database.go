package config

import "fmt"

// DatabaseConfig contains the configuration for the database.
type DatabaseConfig struct {
	Host     string `default:"localhost"`
	Database string `default:"locg_development"`
	Port     string `default:"27017"`
	Username string
	Password string
}

func getConnectionString(username, password, host, port string) string {
	if username != "" && password != "" {
		return fmt.Sprintf("mongodb+srv://%s:%s@%s", username, password, host)
	}
	return fmt.Sprintf("mongodb://%s:%s", host, port)
}

// GetConnectionString returns a string to connect to the database.
func (c *DatabaseConfig) GetConnectionString() string {
	return getConnectionString(c.Username, c.Password, c.Host, c.Port)
}

// GetConnectionStringForDisplay returns a string to connect to the database with password redacted.
func (c *DatabaseConfig) GetConnectionStringForDisplay() string {
	return getConnectionString(c.Username, "********", c.Host, c.Port)
}
