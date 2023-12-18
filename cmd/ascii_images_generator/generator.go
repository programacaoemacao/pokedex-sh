package main

import (
	"fmt"
	"path"

	filehelper "github.com/programacaoemacao/pokedex-sh/app/file_helper"
	progressbar "github.com/programacaoemacao/pokedex-sh/app/gui/progress_bar"
	imggen "github.com/programacaoemacao/pokedex-sh/app/image_generator"
	"github.com/programacaoemacao/pokedex-sh/app/model"
	pathhelper "github.com/programacaoemacao/pokedex-sh/app/path_helper"
)

const (
	pokemonImagesFilename = "pokemon_images.json"
)

func main() {
	var task progressbar.ProbressBarTask = func(inputChannel chan progressbar.ProgressMsg) error {
		repoPath, err := pathhelper.GetRepoRootPath()
		if err != nil {
			return fmt.Errorf("can't get repo path: %v", err)
		}

		inputChannel <- progressbar.ProgressMsg{
			CurrentProgress: float64(0),
			Message:         fmt.Sprintf("Creating %s file", pokemonImagesFilename),
			Type:            progressbar.UpdateProgress,
		}

		// You can change the image generator, check the imggen package
		var imgGenerator imggen.ImageGenerator = imggen.NewDefaultGenerator()
		pokemonImages := [model.LastPokemonID]string{}

		for i := 1; i <= model.LastPokemonID; i++ {
			filepath := fmt.Sprintf("%s/images/%d.png", repoPath, i)
			asciiArt, err := imgGenerator.GenerateAsciiImages(filepath)
			if err != nil {
				return fmt.Errorf("can't convert the Pokémon number %d image", i)
			}

			pokemonImages[i-1] = asciiArt
			inputChannel <- progressbar.ProgressMsg{
				CurrentProgress: float64(i) / float64(model.LastPokemonID),
				Message:         fmt.Sprintf("Converted Pokémon %d to ASCII string", i),
				Type:            progressbar.UpdateProgress,
			}
		}

		inputChannel <- progressbar.ProgressMsg{
			CurrentProgress: float64(1),
			Message:         fmt.Sprintf("Saving Pokémon ASCII data to file %q", pokemonImagesFilename),
			Type:            progressbar.UpdateProgress,
		}

		filePath := path.Join(repoPath, pokemonImagesFilename)
		err = filehelper.CreateFileWithContent(filePath, pokemonImages)
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

	pb := progressbar.NewProgressWriter("Pokémon images to ASCII converter")
	pb.Run(task)
}
