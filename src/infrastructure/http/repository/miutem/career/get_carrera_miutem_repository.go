package career

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	client "github.com/pzentenoe/httpclient-call-go"
	"iluvatar/src/infrastructure/http/repository/miutem/career/entity"
	"io/ioutil"
	"net/http"
)

func (r *carreraMiUTEMRepository) GetCareer(accessToken string) (string, error) {
	headers := http.Header{
		client.HeaderAuthorization: []string{"Bearer " + accessToken},
		echo.HeaderContentType:     []string{echo.MIMEApplicationJSON},
	}

	response, err := r.httpClientCall.
		Path("/v1/carreras/activa").
		Method(http.MethodGet).
		Headers(headers).
		Do()

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("can't get student career, wrong http code from response: %d", response.StatusCode))
	}

	var carreraResponse *entity.CarreraEntity
	err = json.Unmarshal(data, &carreraResponse)
	if err != nil {
		return "", errors.New("can't unmarshal response from server")
	}

	return carreraResponse.Nombre, nil
}
