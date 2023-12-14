package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	imagegenerator "github.com/programacaoemacao/pokedex-sh/app/image_generator"
)

func main() {

	pokemonImages := []string{}
	for i := 1; i <= 1010; i++ {
		filepath := fmt.Sprintf("./images/%d.png", i)
		log.Printf("Converting pokemon image with number %d", i)
		asciiArt, err := imagegenerator.GetImageToPrint(filepath)
		if err != nil {
			log.Fatalf("can't convert the pokemon number %d image", i)
		}
		pokemonImages = append(pokemonImages, asciiArt)
	}

	jsonData, err := json.Marshal(pokemonImages)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	jsonFile, err := os.Create("./pokemon_images.json")
	if err != nil {
		log.Fatalf("Error creating JSON file: %v", err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		log.Fatalf("Error writing JSON data to file: %v", err)
	}
}
