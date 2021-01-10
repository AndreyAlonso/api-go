package middleware

import (
	"github.com/andreyalonso/api-go/clase-7-echo/authorization"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

// Log guarda en los logs la información de cada petición que se realice
func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("petición %q, método: %q", r.URL.Path, r.Method)
		f(w, r)
	}
}

// Authentication valida el token recibido
func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "no permitido"})
		}
		return f(c)
	}
}
