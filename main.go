package main

import (
	"log"

	"github.com/ilopezor/proyecto-twitter-back/bd"
	"github.com/ilopezor/proyecto-twitter-back/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
