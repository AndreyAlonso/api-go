package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/andreyalonso/api-go/clase-7-echo/authorization"
	"github.com/andreyalonso/api-go/clase-7-echo/model"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(c echo.Context) error {
	data := model.Login{}
	err := c.Bind(&data)
	if err != nil {
		resp := newResponse(Error, "estructura no valida", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	if !isLoginValid(&data) {
		resp := newResponse(Error, "usuario o contraseña no validos", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "no se pudo generar el token", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	dataToken := map[string]string{"token": token}
	resp := newResponse(Message, "Login Valido", dataToken)
	return c.JSON(http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contacto@ed.team" && data.Password == "123456"
}
