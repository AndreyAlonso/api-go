package handler

import (
	"github.com/andreyalonso/api-go/clase-3-crud/model"
	"encoding/json"
	"errors"
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

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al crear a la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Persona creada correctamente", nil)
	responseJSON(w, http.StatusCreated, response)
}

// getAll obtener todos los elementos
func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	// Validar que el metodo sea GET
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	resp, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "Hubo un problema al obtener todas las personas", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Ok", resp)
	responseJSON(w, http.StatusOK, response)
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al convertir el slice en json", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
}

func (p *person) getByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "El id debe ser un número entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := p.storage.GetByID(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Ocurrió un error al elminar el registro", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Ok", data)
	responseJSON(w, http.StatusOK, response)
}

// update actualiza el registro seleccionado
func (p *person) update(w http.ResponseWriter, r *http.Request) {
	// Validar que el metodo sea PUT
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	// r.URL.Query() lee todo lo que este en la ruta despues del signo de interrogacion (?)
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "El id debe ser un número entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	// se lee la información de la persona
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al obtener todas las personas", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Persona actualizada correctamente", nil)
	responseJSON(w, http.StatusOK, response)
}

// delete elimina un id obtenido por la url
func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "El id debe  ser un  número entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Ocurrión un error al eliminar el registro", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := newResponse(Message, "Ok", nil)
	responseJSON(w, http.StatusOK, response)
}
