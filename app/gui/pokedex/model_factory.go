package pokedex

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/programacaoemacao/pokedex-sh/app/model"
)

func newModel() (*mainModel, error) {
	m := &mainModel{}

	pokemonJSONFile, err := os.Open("pokemon_info.json")
	if err != nil {
		return nil, fmt.Errorf("error at open the pokemon file")
	}
	defer pokemonJSONFile.Close()

	pokemonImagesFile, err := os.Open("pokemon_images.json")
	if err != nil {
		return nil, fmt.Errorf("error at open the pokemon file")
	}
	defer pokemonImagesFile.Close()

	var pokemons []model.Pokemon
	if err := json.NewDecoder(pokemonJSONFile).Decode(&pokemons); err != nil {
		return nil, fmt.Errorf("error at decoding pokemon json: %+v", err)
	}

	var pokemonImages []string
	if err := json.NewDecoder(pokemonImagesFile).Decode(&pokemonImages); err != nil {
		return nil, fmt.Errorf("error at decoding pokemon images json: %+v", err)
	}

	m.pokemons = pokemons
	m.pokemonImages = pokemonImages

	items := []list.Item{}
	for _, p := range pokemons {
		items = append(items, pokemonInfo{Pokemon: p})
	}

	m.pokedexList = list.New(items, customItemDelegate{}, 18, 44)
	m.pokedexList.Title = "Pokédex"
	m.pokedexList.SetShowPagination(false)

	return m, nil
}
