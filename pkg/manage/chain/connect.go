package chain

import (
	"net/http"

	"buf.build/gen/go/wanho/security-proof-api/connectrpc/go/chain/v1/chainv1connect"
)

// NewChain function is returning a ProofServiceClient accepting a Config.
func NewChain(url string) chainv1connect.ProofServiceClient {
	client := chainv1connect.NewProofServiceClient(
		http.DefaultClient,
		url,
	)

	return client
}
