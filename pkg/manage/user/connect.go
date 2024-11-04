package user

import (
	"net/http"

	"buf.build/gen/go/wanho/security-proof-api/connectrpc/go/api/v1/apiv1connect"
)

// NewUser function is returning a UserServiceClient accepting a Config.
func NewUser(url string) apiv1connect.UserServiceClient {
	client := apiv1connect.NewUserServiceClient(
		http.DefaultClient,
		url,
	)

	return client
}
