package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"

	"github.com/andreyalonso/api-go/clase-7-echo/authorization"
	"github.com/andreyalonso/api-go/clase-7-echo/handler"
	"github.com/andreyalonso/api-go/clase-7-echo/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatal("No se pudo cargar los certificados: %v", err)
	}
	// Inicializar sistema de almacenamiento
	store := storage.NewMemory()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	// Registrar la ruta de persona
	handler.RoutePerson(e, &store)
	handler.RouteLogin(e, &store)

	// Logear que el servidor esta corriendo
	log.Println("Servidor iniciado en el puerto 8080")

	err = e.Start(":8080")
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}

}
