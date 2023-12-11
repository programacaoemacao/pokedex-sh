package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/list"
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
			BorderStyle(lipgloss.RoundedBorder()).
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
	state         sessionState
	pokemonsModel list.Model
	pokemons      []model.Pokemon
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

	m.pokemonsModel = list.New(items, list.NewDefaultDelegate(), 10, 2)
	m.pokemonsModel.Title = "Pokédex"
	return m
}

func (m mainModel) Init() tea.Cmd {
	// start the timer and spinner on program start
	return tea.Batch()
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		print(h, v)
		m.pokemonsModel.SetSize(msg.Width-h, msg.Height-v)
	}

	m.pokemonsModel, _ = m.pokemonsModel.Update(msg)

	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	var s string
	filepath := fmt.Sprintf("./images/%d.png", m.pokemonsModel.Index()+1)
	asciiArt, _ := imagegenerator.GetImageToPrint(filepath)

	currentPokemon := (m.pokemons[m.pokemonsModel.Index()])

	var types string = ""
	switch len(currentPokemon.Types) {
	case 1:
		types = lipgloss.JoinHorizontal(lipgloss.Center,
			pokemonTypeModalStyle.Background(lipgloss.Color(pokemonTypeColors[currentPokemon.Types[0]])).Padding(1, 2).Render(currentPokemon.Types[0]),
		)
	case 2:
		types = lipgloss.JoinHorizontal(lipgloss.Center,
			pokemonTypeModalStyle.Background(lipgloss.Color(pokemonTypeColors[currentPokemon.Types[0]])).MarginRight(1).Padding(1, 2).Render(currentPokemon.Types[0]),
			pokemonTypeModalStyle.Background(lipgloss.Color(pokemonTypeColors[currentPokemon.Types[1]])).MarginLeft(1).Padding(1, 2).Render(currentPokemon.Types[1]),
		)
	}
	imageAndTypeContainer := lipgloss.JoinVertical(lipgloss.Center, pokemonImageModalStyle.Render(asciiArt), types)
	s += lipgloss.JoinHorizontal(lipgloss.Left, listModalStyle.Render(fmt.Sprintf("%4s", m.pokemonsModel.View())), imageAndTypeContainer)
	s += helpStyle.Render("\nq: exit\n")
	return s
}

func main() {
	p := tea.NewProgram(newModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
