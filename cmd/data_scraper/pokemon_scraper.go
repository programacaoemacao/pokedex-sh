package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"sync"

	progressbar "github.com/programacaoemacao/pokedex-sh/app/gui/progress_bar"
	"github.com/programacaoemacao/pokedex-sh/app/model"
	"github.com/programacaoemacao/pokedex-sh/app/scraper"
)

const (
	pokemonsInfoFile = "pokemon_info.json"
)

func main() {
	var task progressbar.ProbressBarTask = func(inputChannel chan progressbar.ProgressMsg) error {
		pokemonList := []*model.Pokemon{}

		currentPath, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		repoPath := strings.Split(currentPath, "cmd")[0]
		repoPath = strings.TrimRight(repoPath, "/")

		const (
			lastPokemonNumber = 1010
		)

		pokeScraper := scraper.NewPokeScraper()
		imageScraper := scraper.NewImageScraper()

		imagesCh := make(chan *model.Pokemon)
		var imagesWG sync.WaitGroup

		go func() {
			for pokemon := range imagesCh {
				err = imageScraper.DownloadImage(pokemon.ImageSrc, fmt.Sprintf("%s/images/%d.png", repoPath, pokemon.ID))
				imagesWG.Done()
				if err != nil {
					// TODO: Handle this error
					log.Fatalf("Error at downloading image: %s", err.Error())
					continue
				}
				inputChannel <- progressbar.ProgressMsg{
					CurrentProgress: float64(pokemon.ID) / float64(lastPokemonNumber),
					Message:         fmt.Sprintf("Downloaded data of Pokémon %s: %q", pokemon.GetFormattedID(), pokemon.Name),
					Type:            progressbar.UpdateProgress,
				}
			}
		}()

		for i := 1; i <= lastPokemonNumber; i++ {
			pokemon, err := pokeScraper.ScrapePokemonInfo(i)
			if err != nil {
				return fmt.Errorf("Error fetching Pokémon %s: %v\n", pokemon.GetFormattedID(), err)
			}
			pokemonList = append(pokemonList, pokemon)

			imagesWG.Add(1)
			imagesCh <- pokemon
		}

		imagesWG.Wait()

		inputChannel <- progressbar.ProgressMsg{
			CurrentProgress: float64(1),
			Message:         fmt.Sprintf("Saving Pokémons data to %s", pokemonsInfoFile),
			Type:            progressbar.UpdateProgress,
		}

		jsonData, err := json.Marshal(pokemonList)
		if err != nil {
			return fmt.Errorf("Error marshaling JSON: %v", err)
		}

		jsonFile, err := os.Create(path.Join(repoPath, pokemonsInfoFile))
		if err != nil {
			return fmt.Errorf("Error creating JSON file: %v", err)
		}
		defer jsonFile.Close()

		_, err = jsonFile.Write(jsonData)
		if err != nil {
			return fmt.Errorf("Error writing JSON data to file: %v", err)
		}

		inputChannel <- progressbar.ProgressMsg{
			CurrentProgress: float64(1),
			Message:         "Done",
			Type:            progressbar.FinishProgram,
		}

		return nil
	}

	pb := progressbar.NewProgressWriter("Scraping Pokémon data")
	pb.Run(task)
}
