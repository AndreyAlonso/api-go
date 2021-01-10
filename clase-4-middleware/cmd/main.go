package main

import (
	"../funciones"
)

func execute(name string, f func(string)) {
	f(name)
}

func main() {
	name := "Comunidad EDteam"
	// Hacer esto no es común
	//saludarConLog := funciones.Log(funciones.Saludar)
	//execute(name, saludarConLog)

	execute(name, funciones.Log(funciones.Saludar))
	execute(name, funciones.Log(funciones.Despedirse))
}
