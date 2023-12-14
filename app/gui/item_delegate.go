package gui

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/programacaoemacao/pokedex-sh/app/model"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type PokemonInfo struct {
	model.Pokemon
}

func (p PokemonInfo) Title() string       { return strconv.Itoa(int(p.ID)) }
func (p PokemonInfo) Description() string { return p.Name }
func (p PokemonInfo) FilterValue() string { return p.Name }

type CustomItemDelegate struct{}

func (d CustomItemDelegate) Height() int                             { return 1 }
func (d CustomItemDelegate) Spacing() int                            { return 0 }
func (d CustomItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d CustomItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	pokemon, ok := listItem.(PokemonInfo)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d - %s", pokemon.ID, pokemon.Name)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
