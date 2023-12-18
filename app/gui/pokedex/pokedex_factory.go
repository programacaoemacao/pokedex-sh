package pokedex

import (
	tea "github.com/charmbracelet/bubbletea"
)

func NewPokedex() (*tea.Program, error) {
	model, err := newModel()
	if err != nil {
		return nil, err
	}

	p := tea.NewProgram(model, tea.WithAltScreen())
	return p, nil
}
