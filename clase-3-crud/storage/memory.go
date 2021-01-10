package storage

import (
	"fmt"
	"github.com/andreyalonso/api-go/clase-3-crud/model"
)

type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

// NewMemory devuelve una instancia de Memory
func NewMemory() Memory {
	persons := make(map[int]model.Person)

	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

// Create
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}
	m.currentID++
	m.Persons[m.currentID] = *person

	return nil
}

// Update actualiza una persona en el slice de memoria
func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	// busca si hay una persona con el ID
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	// La persona si existe con el ID
	m.Persons[ID] = *person

	return nil
}

// Delete borra de la memoria la persona
func (m *Memory) Delete(ID int) error {
	// busca si existe la persona
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	// elimina la persona
	delete(m.Persons, ID)

	return nil
}

// GetByID retorna una persona por el ID
func (m *Memory) GetByID(ID int) (model.Person, error) {
	// se revisa si existe la persona con ID
	person, ok := m.Persons[ID]
	if !ok {
		return person, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}
	return person, nil
}

// GetAll retorna todas las personas que están en la memoria
func (m *Memory) GetAll() (model.Persons, error) {
	var result model.Persons
	// Se recorre el mapa
	for _, v := range m.Persons {
		result = append(result, v)
	}

	return result, nil
}
