package career

import client "github.com/pzentenoe/httpclient-call-go"

type carreraMiUTEMRepository struct {
	httpClientCall *client.HTTPClientCall
}

func NewCareerMiUTEMRepository(httpClientCall *client.HTTPClientCall) *carreraMiUTEMRepository {
	return &carreraMiUTEMRepository{httpClientCall: httpClientCall}
}
