package chain

import (
	"log"

	"github.com/Netflix/go-env"
)

// Config struct composed of a base url.
type Config struct {
	BaseURL string `env:"CHAIN_BASE_URL,default=http://127.0.0.4:8090"`
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
