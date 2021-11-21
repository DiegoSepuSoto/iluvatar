package messaging

import client "github.com/pzentenoe/httpclient-call-go"

type messagingFirebaseRepository struct {
	httpClientCall *client.HTTPClientCall
}

func NewMessagingFirebaseRepository(httpClientCall *client.HTTPClientCall) *messagingFirebaseRepository {
	return &messagingFirebaseRepository{httpClientCall: httpClientCall}
}