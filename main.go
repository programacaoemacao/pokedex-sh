package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	gui "github.com/programacaoemacao/pokedex-sh/app/gui"
	imagegenerator "github.com/programacaoemacao/pokedex-sh/app/image_generator"
	"github.com/programacaoemacao/pokedex-sh/app/model"
)

/*
This example assumes an existing understanding of commands and messages. If you
haven't already read our tutorials on the basics of Bubble Tea and working
with commands, we recommend reading those first.

Find them at:
https://github.com/charmbracelet/bubbletea/tree/master/tutorials/commands
https://github.com/charmbracelet/bubbletea/tree/master/tutorials/basics
*/

var pokemonTypeColors = map[string]string{
	"Normal":   "#A8A77A",
	"Fire":     "#EE8130",
	"Water":    "#6390F0",
	"Electric": "#F7D02C",
	"Grass":    "#7AC74C",
	"Ice":      "#96D9D6",
	"Fighting": "#C22E28",
	"Poison":   "#A33EA1",
	"Ground":   "#E2BF65",
	"Flying":   "#A98FF3",
	"Psychic":  "#F95587",
	"Bug":      "#A6B91A",
	"Rock":     "#B6A136",
	"Ghost":    "#735797",
	"Dragon":   "#6F35FC",
	"Dark":     "#705746",
	"Steel":    "#B7B7CE",
	"Fairy":    "#D685AD",
}

var (
	docStyle       = lipgloss.NewStyle().Margin(2, 2)
	listModalStyle = lipgloss.NewStyle().
			MaxWidth(40).
			Align(lipgloss.Left, lipgloss.Top).
			BorderStyle(lipgloss.NormalBorder()).
			BorderRight(true).
			BorderLeft(false).
			BorderTop(false).
			BorderBottom(false).
			MarginRight(2).
			BorderForeground(lipgloss.Color("69"))
	pokemonImageModalStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.HiddenBorder()).
				BorderForeground(lipgloss.Color("69"))
	pokemonTypeModalStyle = lipgloss.NewStyle().
				BorderForeground(lipgloss.Color("69"))
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type mainModel struct {
	pokedexList list.Model
	pokemons    []model.Pokemon
}

func newModel() mainModel {
	m := mainModel{}

	pokemonJSONFile, err := os.Open("pokemon_info.json")
	if err != nil {
		log.Fatalf("error at open the pokemon file")
	}
	defer pokemonJSONFile.Close()

	var pokemons []model.Pokemon
	if err := json.NewDecoder(pokemonJSONFile).Decode(&pokemons); err != nil {
		log.Fatalf("Error at decodign pokemon json:", err)
	}

	m.pokemons = pokemons

	items := []list.Item{}
	for _, p := range pokemons {
		items = append(items, gui.PokemonInfo{Pokemon: p})
	}

	m.pokedexList = list.New(items, gui.CustomItemDelegate{}, 18, 44)
	m.pokedexList.Title = "Pokédex"
	m.pokedexList.SetShowPagination(false)

	return m
}

func (m mainModel) Init() tea.Cmd {
	// start the timer and spinner on program start
	return tea.Batch()
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	m.pokedexList, cmd = m.pokedexList.Update(msg)
	return m, cmd
}

func (m mainModel) mountPokemonImage(currentPokemon model.Pokemon) string {
	filepath := fmt.Sprintf("./images/%d.png", currentPokemon.ID)
	asciiArt, _ := imagegenerator.GetImageToPrint(filepath)
	return asciiArt
}

func (m mainModel) mountTypes(currentPokemon model.Pokemon) string {

	types := "Types:\t\t"
	blocks := []string{}

	for _, pokemonType := range currentPokemon.Types {
		backgroundColor := pokemonTypeColors[pokemonType]
		block := pokemonTypeModalStyle.
			Background(lipgloss.Color(backgroundColor)).
			MarginRight(1).
			Padding(0, 1).
			Render(pokemonType)
		blocks = append(blocks, block)
	}

	types += strings.Join(blocks, "")
	return types
}

func (m mainModel) mountWeakness(currentPokemon model.Pokemon) string {

	weakness := "Weakness:\t"
	blocks := []string{}

	for _, weakness := range currentPokemon.Weakness {
		backgroundColor := pokemonTypeColors[weakness]
		block := pokemonTypeModalStyle.
			Background(lipgloss.Color(backgroundColor)).
			MarginRight(1).
			Padding(0, 1).
			Render(weakness)
		blocks = append(blocks, block)
	}

	weakness += strings.Join(blocks, "")
	return weakness
}

func breakSentence(sentence string, charactersPerLine int) []string {
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

func (m mainModel) mountAbility(currentPokemon model.Pokemon) string {

	abilitiesText := "Abilities:\n"

	abilities := []string{}
	for i := 0; i < len(currentPokemon.Abilities); i++ {
		ability := currentPokemon.Abilities[i]
		abilityInfo := currentPokemon.AbilitiesInfo[i]
		abilityText := fmt.Sprintf("%d. %s: %s", i+1, ability, abilityInfo)
		abilities = append(abilities, breakSentence(abilityText, 70)...)
	}

	abilitiesText += strings.Join(abilities, "\n")
	return abilitiesText
}

func (m mainModel) mountHeightWeight(currentPokemon model.Pokemon) string {
	heightWeight := fmt.Sprintf("Height: %s\tWeight: %s", currentPokemon.Height, currentPokemon.Weight)
	return heightWeight
}

func (m mainModel) mountCategory(currentPokemon model.Pokemon) string {
	category := fmt.Sprintf("Category:\t%s", currentPokemon.Category)
	return category
}

func (m mainModel) View() string {
	currentPokemon := m.pokedexList.SelectedItem().(gui.PokemonInfo).Pokemon

	var s string
	asciiArt := m.mountPokemonImage(currentPokemon)
	category := m.mountCategory(currentPokemon)
	types := m.mountTypes(currentPokemon)
	weakness := m.mountWeakness(currentPokemon)
	heightWeight := m.mountHeightWeight(currentPokemon)
	ability := m.mountAbility(currentPokemon)

	imageAndTypeContainer := lipgloss.JoinVertical(lipgloss.Left, pokemonImageModalStyle.Render(asciiArt), category, types, weakness, heightWeight, ability)
	s += lipgloss.JoinHorizontal(lipgloss.Left, m.pokedexList.View(), imageAndTypeContainer)
	s += helpStyle.Render("\nq: exit\n")
	return s
}

func main() {
	p := tea.NewProgram(newModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
