package handler

import (
	"errors"
	"github.com/andreyalonso/api-go/clase-7-echo/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// Person estructura que tiene un storage
type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(c echo.Context) error {
	data := model.Person{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al crear a la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Persona creada correctamente", nil)
	return c.JSON(http.StatusCreated, response)
}

// getAll obtener todos los elementos
func (p *person) getAll(c echo.Context) error {
	resp, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "Hubo un problema al obtener todas las personas", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Ok", resp)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getByID(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "El id debe ser un número entero positivo", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data, err := p.storage.GetByID(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		return c.JSON(http.StatusBadRequest, response)

	}
	if err != nil {
		response := newResponse(Error, "Ocurrió un error al consultar el registro", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Ok", data)
	return c.JSON(http.StatusOK, response)
}

// update actualiza el registro seleccionado
func (p *person) update(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "El id debe ser un número entero positivo", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := model.Person{}
	err = c.Bind(&data) // se lee la información de la persona
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al obtener todas las personas", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Persona actualizada correctamente", nil)
	return c.JSON(http.StatusOK, response)
}

// delete elimina un id obtenido por la url
func (p *person) delete(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "El id debe  ser un  número entero positivo", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		return c.JSON(http.StatusInternalServerError, response)

	}
	if err != nil {
		response := newResponse(Error, "Ocurrió un error al eliminar el registro", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := newResponse(Message, "Ok", nil)
	return c.JSON(http.StatusOK, response)
}
