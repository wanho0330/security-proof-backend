// Package db is a package for handling database processes.
package db

import (
	"log"
	"strings"

	"github.com/Netflix/go-env"
)

// WriteConfig struct is composed of a host, a port, a ssl mode, a name, a schema, an id and passwd.
type WriteConfig struct {
	Host    string `env:"DB_WRITE_HOST,default=localhost"`
	Port    string `env:"DB_WRITE_PORT,default=5432"`
	SSLMode string `env:"DB_WRITE_SSL_MODE,default=disable"`
	Name    string `env:"DB_WRITE_NAME,default=security_proof"`
	Schema  string `env:"DB_WRITE_SCHEMA,default=user"`
	ID      string `env:"DB_WRITE_ID,default=postgres"`
	Passwd  string `env:"DB_WRITE_PASSWD,default=postgres"`
}

// ReadConfig struct is composed of a host, a port, a ssl mode, a name, a schema, an id and passwd.
type ReadConfig struct {
	Host    string `env:"DB_Read_HOST,default=localhost"`
	Port    string `env:"DB_Read_PORT,default=5432"`
	SSLMode string `env:"DB_Read_SSL_MODE,default=disable"`
	Name    string `env:"DB_Read_NAME,default=security_proof"`
	Schema  string `env:"DB_Read_SCHEMA,default=user"`
	ID      string `env:"DB_Read_ID,default=postgres"`
	Passwd  string `env:"DB_Read_PASSWD,default=postgres"`
}

// TokenConfig struct is composed of a host, a port, an id, a passwd and schema.
type TokenConfig struct {
	Host   string `env:"DB_TOKEN_HOST,default=localhost"`
	Port   string `env:"DB_TOKEN_PORT,default=6379"`
	ID     string `env:"DB_TOKEN_ID,default="`
	Passwd string `env:"DB_TOKEN_PASSWD,default="`
	Schema string `env:"DB_TOKEN_SCHEMA,default=0"`
}

// WriteConfig method is returning a WriteConfig.
func (c *WriteConfig) WriteConfig() *WriteConfig {
	return c
}

// Dsn method is returning a dsn string.
func (c *WriteConfig) Dsn() string {
	_, err := env.UnmarshalFromEnviron(c)
	if err != nil {
		log.Fatal(err)
	}

	if c == nil {
		return ""
	}

	dsn := strings.Builder{}
	dsn.WriteString("postgresql://")
	dsn.WriteString(c.ID)
	dsn.WriteString(":")
	dsn.WriteString(c.Passwd)
	dsn.WriteString("@")
	dsn.WriteString(c.Host)
	dsn.WriteString(":")
	dsn.WriteString(c.Port)
	dsn.WriteString("/")
	dsn.WriteString(c.Name)
	dsn.WriteString("?sslmode=")
	dsn.WriteString(c.SSLMode)

	return dsn.String()
}

// ReadConfig method is a ReadConfig.
func (c *ReadConfig) ReadConfig() *ReadConfig {
	return c
}

// Dsn method is returning a dsn string.
func (c *ReadConfig) Dsn() string {
	_, err := env.UnmarshalFromEnviron(c)
	if err != nil {
		log.Fatal(err)
	}

	if c == nil {
		return ""
	}
	dsn := strings.Builder{}
	dsn.WriteString("postgresql://")
	dsn.WriteString(c.ID)
	dsn.WriteString(":")
	dsn.WriteString(c.Passwd)
	dsn.WriteString("@")
	dsn.WriteString(c.Host)
	dsn.WriteString(":")
	dsn.WriteString(c.Port)
	dsn.WriteString("/")
	dsn.WriteString(c.Name)
	dsn.WriteString("?sslmode=")
	dsn.WriteString(c.SSLMode)

	return dsn.String()
}

// TokenConfig method is a returning a TokenConfig.
func (c *TokenConfig) TokenConfig() *TokenConfig { return c }

// Dsn method is a returning dsn string.
func (c *TokenConfig) Dsn() string {
	if c == nil {
		return ""
	}

	dsn := strings.Builder{}
	dsn.WriteString("redis://")
	dsn.WriteString(c.ID)
	dsn.WriteString(":")
	dsn.WriteString(c.Passwd)
	dsn.WriteString("@")
	dsn.WriteString(c.Host)
	dsn.WriteString(":")
	dsn.WriteString(c.Port)
	dsn.WriteString("/")
	dsn.WriteString(c.Schema)

	return dsn.String()
}
