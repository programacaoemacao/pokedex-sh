package main

import (
	"log"

	guipokedex "github.com/programacaoemacao/pokedex-sh/app/gui/pokedex"
)

func main() {
	program, err := guipokedex.NewPokedex()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := program.Run(); err != nil {
		log.Fatal(err)
	}
}
