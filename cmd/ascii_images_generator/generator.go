package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	imggen "github.com/programacaoemacao/pokedex-sh/app/image_generator"
	"github.com/programacaoemacao/pokedex-sh/app/model"
	pathhelper "github.com/programacaoemacao/pokedex-sh/app/path_helper"
)

func main() {
	repoPath, err := pathhelper.GetRepoRootPath()
	if err != nil {
		log.Fatalf("can't get repo path: %v", err)
	}

	pokemonImgsJSON, err := os.Create(fmt.Sprintf("%s/pokemon_images.json", repoPath))
	if err != nil {
		log.Fatalf("Error creating JSON file: %v", err)
	}
	defer pokemonImgsJSON.Close()

	// You can change the image generator
	var imgGenerator imggen.ImageGenerator = imggen.NewDefaultGenerator()

	pokemonImages := [model.LastPokemonID]string{}
	for i := 1; i <= model.LastPokemonID; i++ {
		filepath := fmt.Sprintf("%s/images/%d.png", repoPath, i)
		log.Printf("Converting pokemon image with number %d", i)
		asciiArt, err := imgGenerator.GenerateAsciiImages(filepath)
		if err != nil {
			log.Fatalf("can't convert the pokemon number %d image", i)
		}
		pokemonImages[i-1] = asciiArt
	}

	jsonData, err := json.Marshal(pokemonImages)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	_, err = pokemonImgsJSON.Write(jsonData)
	if err != nil {
		log.Fatalf("Error writing JSON data to file: %v", err)
	}
}
