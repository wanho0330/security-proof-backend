// Package user is a package for handling user processes.
package user

import (
	"log"

	"github.com/Netflix/go-env"
)

// Config struct composed of a base url.
type Config struct {
	BaseURL string `env:"USER_BASE_URL,default=http://127.0.0.1:8080"`
}

// FromEnv function is returning base url.
func (c *Config) FromEnv() string {
	_, err := env.UnmarshalFromEnviron(c)
	if err != nil {
		log.Fatal("Error unmarshalling environment variables")
		return ""
	}
	return c.BaseURL
}
