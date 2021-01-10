// En handler va a crear la interface que debe implementar los sistemas de almacenamiento de la api

package handler

import "github.com/andreyalonso/api-go/clase-7-echo/model"

// Storage
type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	GetByID(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}
