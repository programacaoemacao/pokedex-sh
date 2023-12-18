package pokedex

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/programacaoemacao/pokedex-sh/app/model"
)

type mainModel struct {
	pokedexList   list.Model
	pokemons      []model.Pokemon
	pokemonImages []string
}

func (m *mainModel) Init() tea.Cmd {
	return tea.Batch()
}

func (m *mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.pokedexList, cmd = m.pokedexList.Update(msg)
	return m, cmd
}

func (m *mainModel) mountPokemonImage(currentPokemon model.Pokemon) string {
	asciiArt := m.pokemonImages[currentPokemon.ID-1]
	return asciiArt
}

func (m *mainModel) mountTypes(currentPokemon model.Pokemon) string {
	types := "Types:   \t"
	blocks := []string{}

	for _, pokemonType := range currentPokemon.Types {
		backgroundColor := pokemonTypeColors[pokemonType]
		block := pokemonTypeModalStyle.
			Background(lipgloss.Color(backgroundColor)).
			Render(pokemonType)
		blocks = append(blocks, block)
	}

	types += strings.Join(blocks, "")
	return types
}

func (m *mainModel) mountWeakness(currentPokemon model.Pokemon) string {
	weakness := "Weakness:\t"
	blocks := []string{}

	for _, weakness := range currentPokemon.Weakness {
		backgroundColor := pokemonTypeColors[weakness]
		block := pokemonTypeModalStyle.
			Background(lipgloss.Color(backgroundColor)).
			Render(weakness)
		blocks = append(blocks, block)
	}

	weakness += strings.Join(blocks, "")
	return weakness
}

func (m *mainModel) breakSentence(sentence string, charactersPerLine int) []string {
	words := strings.Fields(sentence)
	var lines []string
	var currentLine string

	for _, word := range words {
		if len(currentLine)+len(word)+1 <= charactersPerLine {
			if currentLine == "" {
				currentLine = word
			} else {
				currentLine += " " + word
			}
		} else {
			lines = append(lines, currentLine)
			currentLine = word
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

func (m *mainModel) mountAbility(currentPokemon model.Pokemon) string {
	abilitiesText := "Abilities:\n"

	abilities := []string{}
	for i := 0; i < len(currentPokemon.Abilities); i++ {
		ability := currentPokemon.Abilities[i]
		abilityInfo := currentPokemon.AbilitiesInfo[i]
		abilityText := fmt.Sprintf("%d. %s: %s", i+1, ability, abilityInfo)
		abilities = append(abilities, m.breakSentence(abilityText, 70)...)
	}

	abilitiesText += strings.Join(abilities, "\n")
	return abilitiesText
}

func (m *mainModel) mountHeightWeight(currentPokemon model.Pokemon) string {
	heightWeight := fmt.Sprintf("Height: %s\tWeight: %s", currentPokemon.Height, currentPokemon.Weight)
	return heightWeight
}

func (m *mainModel) mountCategory(currentPokemon model.Pokemon) string {
	category := fmt.Sprintf("Category:\t%s", currentPokemon.Category)
	return category
}

func (m *mainModel) View() string {
	currentPokemon := m.pokedexList.SelectedItem().(pokemonInfo).Pokemon
	asciiArt := m.mountPokemonImage(currentPokemon)
	category := m.mountCategory(currentPokemon)
	types := m.mountTypes(currentPokemon)
	weakness := m.mountWeakness(currentPokemon)
	heightWeight := m.mountHeightWeight(currentPokemon)
	ability := m.mountAbility(currentPokemon)

	var s string
	imageAndTypeContainer := lipgloss.JoinVertical(lipgloss.Left, pokemonImageModalStyle.Render(asciiArt), category, types, weakness, heightWeight, ability)
	s += lipgloss.JoinHorizontal(lipgloss.Left, m.pokedexList.View(), imageAndTypeContainer)
	return s
}
