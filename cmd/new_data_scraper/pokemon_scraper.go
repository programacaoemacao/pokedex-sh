package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/programacaoemacao/pokedex-sh/app/model"
	"github.com/programacaoemacao/pokedex-sh/app/scraper"
)

func main() {
	pokemonList := []*model.Pokemon{}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	repoPath := strings.Split(path, "cmd")[0]
	repoPath = strings.TrimRight(repoPath, "/")

	const (
		lastPokemonNumber = 1010
	)

	pokeScraper := scraper.NewPokeScraper()
	imageScraper := scraper.NewImageScraper()

	imagesCh := make(chan *model.Pokemon)
	var imagesWG sync.WaitGroup

	go func() {
		log.Printf("Initializing the images download goroutine")
		for pokemon := range imagesCh {
			log.Printf("Downloading image for: %s\n", pokemon.Name)
			err = imageScraper.DownloadImage(pokemon.ImageSrc, fmt.Sprintf("%s/images/%d.png", repoPath, pokemon.ID))
			if err != nil {
				log.Printf("Error at downloading image: %s", err.Error())
				return
			}
			log.Printf("Downloaded image for: %s\n", pokemon.Name)
			imagesWG.Done()
		}
	}()

	for i := 1; i <= lastPokemonNumber; i++ {
		pokemon, err := pokeScraper.ScrapePokemonInfo(i)
		if err != nil {
			log.Printf("Error fetching Pokémon information with ID %d: %v\n", i, err)
			continue
		}
		log.Printf("Pokemon information with ID %d:\n", pokemon.ID)
		log.Printf("Name: %s\n", pokemon.Name)
		log.Printf("Height: %s\n", pokemon.Height)
		log.Printf("Weight: %s\n", pokemon.Weight)
		log.Printf("Category: %s\n", pokemon.Category)
		log.Printf("Abilities: %v\n", pokemon.Abilities)
		log.Printf("Types: %v\n", pokemon.Types)
		log.Printf("Weaknesses: %v\n", pokemon.Weakness)
		log.Printf("Image Source: %s\n", pokemon.ImageSrc)
		fmt.Println("---------------------------------------")
		pokemonList = append(pokemonList, pokemon)

		imagesWG.Add(1)
		imagesCh <- pokemon
	}

	imagesWG.Wait()

	jsonData, err := json.MarshalIndent(pokemonList, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	jsonFile, err := os.Create("pokemon_info.json")
	if err != nil {
		log.Fatalf("Error creating JSON file: %v", err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		log.Fatalf("Error writing JSON data to file: %v", err)
	}

	fmt.Println("Pokemon information has been saved to 'pokemon_info.json'")
}
