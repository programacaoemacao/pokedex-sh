package pokedex

import "github.com/charmbracelet/lipgloss"

var (
	docStyle = lipgloss.NewStyle().Margin(2, 2)

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
				BorderForeground(lipgloss.Color("69")).ColorWhitespace(false)

	pokemonTypeModalStyle = lipgloss.NewStyle().
				BorderForeground(lipgloss.Color("69")).
				MarginRight(1).
				Padding(0, 1)
)
