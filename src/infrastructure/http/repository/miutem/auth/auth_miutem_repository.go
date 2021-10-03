package auth

import (
	client "github.com/pzentenoe/httpclient-call-go"
)

type authMiUTEMRepository struct {
	httpClientCall *client.HTTPClientCall
}

func NewLoginMiUTEMRepository(httpClientCall *client.HTTPClientCall) *authMiUTEMRepository {
	return &authMiUTEMRepository{httpClientCall: httpClientCall}
}
