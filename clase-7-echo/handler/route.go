package handler

import (
	"github.com/andreyalonso/api-go/clase-7-echo/middleware"
	"github.com/labstack/echo"
)

// RoutePerson
func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)
	person := e.Group("/v1/persons")
	person.Use(middleware.Authentication)

	person.POST("", h.create)
	person.GET("", h.getAll)
	person.GET("/:id", h.getByID)
	person.PUT("/:id", h.update)
	person.DELETE("/:id", h.delete)
}

// RouteLogin rutas que necesitan un inicio de sesión
func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)

	e.POST("/v1/login", h.login)
}
