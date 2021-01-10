package model

// Community estructura de una comunidad
type Community struct {
	// Name nombre de una comunidad. Ejemplo: EDteam
	Name string `json:"name"`
}

// Communities slice de comunidades
type Communities []Community

// Person estructura de una persona
type Person struct {
	ID uint `json:"id"`
	// Name nombre de la persona
	Name string `json:"name"`
	// Age edad de la persona
	Age uint8 `json:"age"`
	// Communities comunidades a las que pertenece una persona
	Communities Communities `json:"communities"`
}

// Persons slice de personas
type Persons []Person
