package elastic

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

type elastic struct {
	client *elasticsearch.TypedClient
}

// NewElastic function is returning an Elastic interface, accepting a Config.
func NewElastic(addresses []string, username string, password string, caCert []byte) Elastic {
	// 내부망이므로 인증서에 대한 유효성 검증 off
	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: addresses,
		Username:  username,
		Password:  password,
		CACert:    caCert,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return &elastic{client: client}
}
