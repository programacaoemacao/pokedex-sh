package pokedex

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

type pokemonInfo struct {
	model.Pokemon
}

func (p pokemonInfo) Title() string {
	return strconv.Itoa(int(p.ID))
}

func (p pokemonInfo) Description() string {
	return p.Name
}
func (p pokemonInfo) FilterValue() string {
	return p.Name
}

type customItemDelegate struct{}

func (d customItemDelegate) Height() int {
	return 1
}

func (d customItemDelegate) Spacing() int {
	return 0
}

func (d customItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

func (d customItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	pokemon, ok := listItem.(pokemonInfo)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d - %s", pokemon.ID, pokemon.Name)

	renderFn := itemStyle.Render
	if index == m.Index() {
		renderFn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, renderFn(str))
}
