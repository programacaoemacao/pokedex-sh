package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

// sessionState is used to track which model is focused
type sessionState uint

const (
	listView sessionState = iota
	spinnerView
)

type pokemonInfo struct {
	model.Pokemon
}

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

func (p pokemonInfo) Title() string       { return strconv.Itoa(int(p.ID)) }
func (p pokemonInfo) Description() string { return p.Name }
func (p pokemonInfo) FilterValue() string { return p.Name }

var (
	docStyle       = lipgloss.NewStyle().Margin(2, 2)
	listModalStyle = lipgloss.NewStyle().
			MaxWidth(50).
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
	state        sessionState
	pokedexTable table.Model
	pokemons     []model.Pokemon
}

func newModel() mainModel {
	m := mainModel{state: listView}

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
		items = append(items, pokemonInfo{p})
	}

	columns := []table.Column{
		{Title: "ID", Width: 4},
		{Title: "Name", Width: 14},
	}

	rows := []table.Row{}
	for _, pokemon := range pokemons {
		row := table.Row{
			strconv.Itoa(pokemon.ID),
			pokemon.Name,
		}
		rows = append(rows, row)
	}

	m.pokedexTable = table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(39),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	m.pokedexTable.SetStyles(s)

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

	m.pokedexTable, cmd = m.pokedexTable.Update(msg)

	return m, cmd
}

func (m mainModel) mountPokemonImage(pokemonID int) string {
	filepath := fmt.Sprintf("./images/%d.png", pokemonID)
	asciiArt, _ := imagegenerator.GetImageToPrint(filepath)
	return asciiArt
}

func (m mainModel) mountTypes(pokemonID int) string {
	currentPokemon := m.pokemons[pokemonID-1]

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

func (m mainModel) mountWeakness(pokemonID int) string {
	currentPokemon := m.pokemons[pokemonID-1]

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

func (m mainModel) mountAbility(pokemonID int) string {
	currentPokemon := m.pokemons[pokemonID-1]

	abilitiesText := "Abilities:\n"

	abilities := []string{}
	for i := 0; i < len(currentPokemon.Abilities); i++ {
		ability := currentPokemon.Abilities[i]
		abilityInfo := currentPokemon.AbilitiesInfo[i]
		abilityText := fmt.Sprintf("\t%s: %s", ability, abilityInfo)
		abilities = append(abilities, abilityText)
	}

	abilitiesText += strings.Join(abilities, "\n")
	if len(currentPokemon.Abilities) == 1 {
		abilitiesText += "\n"
	}
	return abilitiesText
}

func (m mainModel) mountHeightWeight(pokemonID int) string {
	currentPokemon := m.pokemons[pokemonID-1]
	heightWeight := fmt.Sprintf("Height:%s\tWeight:%s", currentPokemon.Height, currentPokemon.Weight)
	return heightWeight
}

func (m mainModel) View() string {
	stringID := m.pokedexTable.SelectedRow()[0]
	id, _ := strconv.Atoi(stringID)

	var s string
	asciiArt := m.mountPokemonImage(id)
	types := m.mountTypes(id)
	weakness := m.mountWeakness(id)
	ability := m.mountAbility(id)
	heightWeight := m.mountHeightWeight(id)

	imageAndTypeContainer := lipgloss.JoinVertical(lipgloss.Left, pokemonImageModalStyle.Render(asciiArt), types, "", weakness, "", ability, "", heightWeight)
	s += lipgloss.JoinHorizontal(lipgloss.Left, listModalStyle.Render(fmt.Sprintf("%4s", m.pokedexTable.View())), imageAndTypeContainer)
	s += helpStyle.Render("\nq: exit\n")
	return s
}

func main() {
	p := tea.NewProgram(newModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
