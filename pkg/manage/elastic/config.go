// Package elastic is a package for handling elastic processes.
package elastic

import (
	"log"
	"os"

	"github.com/Netflix/go-env"
)

// Config struct is composed of an address, a username, a password and ca cert path.
type Config struct {
	Address    string `env:"ELA_ADDRESS,default=https://127.0.0.1:9200"`
	UserName   string `env:"ELA_USERNAME,default=elastic"`
	Password   string `env:"ELA_PASSWORD,default=Ty9-rHwUHi=4kBUIdwuM"`
	CACertPath string `env:"ELA_CA_CERT_PATH,default=/Users/wanho/GolandProjects/security-proof/config/certs/http_ca.crt"`
}

// FromEnv method is returning addresses, username, password, cert.
func (c *Config) FromEnv() (addresses []string, username string, password string, caCert []byte) {
	_, err := env.UnmarshalFromEnviron(c)
	if err != nil {
		log.Fatal("Error unmarshalling environment variables")
		return nil, "", "", nil
	}

	cert, err := os.ReadFile(c.CACertPath)
	if err != nil {
		log.Fatal("Error reading CA certificate")
	}

	return []string{c.Address}, c.UserName, c.Password, cert
}

// Index constants is elastic search index.
var Index = "fluentd"
