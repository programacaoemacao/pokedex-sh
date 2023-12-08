package scraper

import (
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gocolly/colly"
	"github.com/programacaoemacao/pokedex-sh/app/model"
	"github.com/stretchr/testify/require"
)

func TestMountURLToScrape(t *testing.T) {
	scraper := pokeScraper{}

	t.Run("Run with number 1", func(t *testing.T) {
		actualURL := scraper.mountURLToScrape(1)
		expectedURL := "https://sg.portal-pokemon.com/play/pokedex/0001"
		require.Equal(t, actualURL, expectedURL)
	})

	t.Run("Run with number 91", func(t *testing.T) {
		actualURL := scraper.mountURLToScrape(91)
		expectedURL := "https://sg.portal-pokemon.com/play/pokedex/0091"
		require.Equal(t, actualURL, expectedURL)
	})

	t.Run("Run with number 1000", func(t *testing.T) {
		actualURL := scraper.mountURLToScrape(1000)
		expectedURL := "https://sg.portal-pokemon.com/play/pokedex/1000"
		require.Equal(t, actualURL, expectedURL)
	})
}

func TestScrape(t *testing.T) {
	absPath := getAbsoluteProjectRootDir(t)
	s := createNewFileScraper()

	t.Run("Collect Kilowattrel data", func(t *testing.T) {
		url := "file://" + absPath + "/test_htmls/pokemon_0941.html"
		pokemon, err := s.scrape(url)
		require.NoError(t, err)

		expectedPokemon := &model.Pokemon{
			ID:     941,
			Name:   "Kilowattrel",
			Height: "1.4 m",
			Weight: "38.6 kg",
			Abilities: []string{
				"Wind Power",
				"Volt Absorb",
			},
			AbilitiesInfo: []string{
				"The Pokémon becomes charged when it is hit by a wind move, boosting the power of the next Electric-type move the Pokémon uses.",
				"Restores HP if hit by an Electric-type move, instead of taking damage.",
			},
			Category: "Frigatebird Pokémon",
			Types:    []string{"Electric", "Flying"},
			Weakness: []string{"Ice", "Rock"},
			ImageSrc: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/941.png",
		}

		require.Equal(t, pokemon, expectedPokemon)
	})
}

func createNewFileScraper() pokeScraper {
	scraper := pokeScraper{}
	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c := colly.NewCollector()
	c.WithTransport(transport)

	scraper.collector = c

	return scraper
}

func getAbsoluteProjectRootDir(t *testing.T) string {
	dir, err := filepath.Abs(".")
	require.NoError(t, err)
	return strings.Split(dir, "pokedex-sh")[0] + "pokedex-sh"
}
