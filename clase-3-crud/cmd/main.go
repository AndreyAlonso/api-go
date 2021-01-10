package main

import (
	"github.com/andreyalonso/api-go/clase-3-crud/authorization"
	"github.com/andreyalonso/api-go/clase-3-crud/handler"
	"github.com/andreyalonso/api-go/clase-3-crud/storage"
	"log"
	"net/http"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatal("No se pudo cargar los certificados: %v", err)
	}
	// Inicializar sistema de almacenamiento
	store := storage.NewMemory()
	// Inicializar server mux
	mux := http.NewServeMux()
	// Registrar la ruta de persona
	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	// Logear que el servidor esta corriendo
	log.Println("Servidor iniciado en el puerto 8080")

	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}

}
