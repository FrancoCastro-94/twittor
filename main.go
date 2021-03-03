package main

import (
	"log"

	"github.com/FrancoCastro-94/twittor/bd"
	"github.com/FrancoCastro-94/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
