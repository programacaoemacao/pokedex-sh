package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

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

	fileContent, err := os.ReadFile(fmt.Sprintf("%s/pokemon_data_input.json", repoPath))
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	type pokemonInfo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	pokemonsInfo := []pokemonInfo{}
	err = json.Unmarshal(fileContent, &pokemonsInfo)
	if err != nil {
		fmt.Println("Erro ao fazer unmarshal do JSON:", err)
		return
	}

	pokeScraper := scraper.NewPokeScraper()

	for _, info := range pokemonsInfo {
		pokemonInfo, err := pokeScraper.ScrapePokemonInfo(info.ID, info.Name)
		if err != nil {
			log.Printf("Error fetching Pokémon information with ID %d: %v\n", info.ID, err)
			continue
		}
		fmt.Printf("Pokemon information with ID %d:\n", info.ID)
		fmt.Printf("Name: %s\n", pokemonInfo.Name)
		fmt.Printf("Height: %s\n", pokemonInfo.Height)
		fmt.Printf("Weight: %s\n", pokemonInfo.Weight)
		fmt.Printf("Category: %s\n", pokemonInfo.Category)
		fmt.Printf("Abilities: %v\n", pokemonInfo.Abilities)
		fmt.Printf("Types: %v\n", pokemonInfo.Types)
		fmt.Printf("Weaknesses: %v\n", pokemonInfo.Weakness)
		fmt.Printf("Image Source: %s\n", pokemonInfo.ImageSrc)
		fmt.Println("---------------------------------------")
		pokemonList = append(pokemonList, pokemonInfo)
	}

	for _, pokemon := range pokemonList {
		fmt.Printf("Downloading image for %s\n", pokemon.Name)
		err = scraper.DownloadImage(pokemon.ImageSrc, fmt.Sprintf("%s/images/%d.png", repoPath, pokemon.ID))
		if err != nil {
			fmt.Printf("Error at downloading image: %s", err.Error())
			return
		}
		fmt.Printf("Downloaded image!\n")
	}

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
