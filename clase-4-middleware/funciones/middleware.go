package funciones

import (
	"fmt"
	"time"
)

type MyFunction func(string)

// Log .
func Log(f MyFunction) MyFunction {
	return func(name string) {
		fmt.Println("Inicio: ", time.Now().Format("2006-01-02 15:04:05"))
		f(name)
		fmt.Println("Fin: ", time.Now().Format("2006-01-02 15:04:05"), "\n")
	}
}
