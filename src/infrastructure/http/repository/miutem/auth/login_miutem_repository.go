package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"iluvatar/src/domain/models"
	"iluvatar/src/infrastructure/http/repository/miutem/auth/entity"
	"iluvatar/src/infrastructure/http/repository/miutem/auth/mapper"
	"io/ioutil"
	"net/http"
)

func (r *authMiUTEMRepository) Login(email, password string) (*models.Student, string, error) {
	body := echo.Map{"usuario": email, "contrasenia": password}
	headers := http.Header{
		echo.HeaderContentType: []string{echo.MIMEApplicationJSON},
	}

	response, err := r.httpClientCall.
		Path("/v1/usuarios/login").
		Method(http.MethodPost).
		Headers(headers).
		Body(body).
		Do()

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, "", err
	}

	if response.StatusCode != http.StatusOK {
		return nil, "", errors.New(fmt.Sprintf("can't login user, wrong http code from response: %d", response.StatusCode))
	}

	var loginResponse *entity.LoginEntity
	err = json.Unmarshal(data, &loginResponse)
	if err != nil {
		return nil, "", errors.New("can't unmarshal response from server")
	}

	if isStudent(loginResponse) {
		return mapper.MapLoginEntityToStudentModel(loginResponse), loginResponse.Sesion, nil
	} else {
		return nil, "", errors.New("user is not an student")
	}
}

func isStudent(loginEntity *entity.LoginEntity) bool {
	for _, userType := range loginEntity.TiposUsuario {
		if userType == "Estudiante" {
			return true
		}
	}

	return false
}