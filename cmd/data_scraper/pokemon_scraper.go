package main

import (
	"fmt"
	"log"
	"path"
	"sync"

	filehelper "github.com/programacaoemacao/pokedex-sh/app/file_helper"
	progressbar "github.com/programacaoemacao/pokedex-sh/app/gui/progress_bar"
	"github.com/programacaoemacao/pokedex-sh/app/model"
	pathhelper "github.com/programacaoemacao/pokedex-sh/app/path_helper"
	"github.com/programacaoemacao/pokedex-sh/app/scraper"
)

const (
	pokemonsInfoFile = "pokemon_info.json"
)

func main() {
	var task progressbar.ProbressBarTask = func(inputChannel chan progressbar.ProgressMsg) error {
		pokemonList := []*model.Pokemon{}

		repoPath, err := pathhelper.GetRepoRootPath()
		if err != nil {
			return err
		}

		pokeScraper := scraper.NewPokeScraper()
		imageScraper := scraper.NewImageScraper()

		imagesCh := make(chan *model.Pokemon)
		var imageDownloadWG sync.WaitGroup

		go func() {
			for pokemon := range imagesCh {
				err = imageScraper.DownloadImage(pokemon.ImageSrc, fmt.Sprintf("%s/images/%d.png", repoPath, pokemon.ID))
				imageDownloadWG.Done()
				if err != nil {
					// TODO: Handle this error
					log.Fatalf("Error at downloading image: %s", err.Error())
					continue
				}
				inputChannel <- progressbar.ProgressMsg{
					CurrentProgress: float64(pokemon.ID) / float64(model.LastPokemonID),
					Message:         fmt.Sprintf("Scraped data from Pokémon number %s: %q", pokemon.GetFormattedID(), pokemon.Name),
					Type:            progressbar.UpdateProgress,
				}
			}
		}()

		for i := 1; i <= model.LastPokemonID; i++ {
			pokemon, err := pokeScraper.ScrapePokemonInfo(i)
			if err != nil {
				return fmt.Errorf("Error fetching Pokémon %s: %v\n", pokemon.GetFormattedID(), err)
			}
			pokemonList = append(pokemonList, pokemon)

			imageDownloadWG.Add(1)
			imagesCh <- pokemon
		}

		imageDownloadWG.Wait()

		inputChannel <- progressbar.ProgressMsg{
			CurrentProgress: float64(1),
			Message:         fmt.Sprintf("Saving Pokémons data to %s", pokemonsInfoFile),
			Type:            progressbar.UpdateProgress,
		}

		filePath := path.Join(repoPath, pokemonsInfoFile)
		err = filehelper.CreateFileWithContent(filePath, pokemonList)
		if err != nil {
			return fmt.Errorf("error at saving data to json: %+v", err)
		}

		inputChannel <- progressbar.ProgressMsg{
			CurrentProgress: float64(1),
			Message:         "Done",
			Type:            progressbar.FinishProgram,
		}

		return nil
	}

	pb := progressbar.NewProgressWriter("Pokémon data Scraper")
	pb.Run(task)
}
