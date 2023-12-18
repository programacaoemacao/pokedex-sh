package progressbar

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

const (
	padding  = 2
	maxWidth = 80
)

type messageType int

const (
	UpdateProgress messageType = iota
	FinishProgram  messageType = iota
)

type ProgressMsg struct {
	CurrentProgress float64
	Message         string
	Type            messageType
}
type ProgressErrMsg struct{ err error }

func finalPause() tea.Cmd {
	return tea.Tick(time.Millisecond*750, func(_ time.Time) tea.Msg {
		return nil
	})
}

type progressModel struct {
	progress   progress.Model
	err        error
	taskTitle  string
	currentLog string
}

func (m progressModel) Init() tea.Cmd {
	return nil
}

func (m progressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case ProgressErrMsg:
		m.err = msg.err
		return m, tea.Quit

	case ProgressMsg:
		var cmds []tea.Cmd
		m.progress.SetPercent(float64(msg.CurrentProgress))
		progressModel, updateCmd := m.progress.Update(msg)
		cmds = append(cmds, updateCmd)
		if msg.Type == FinishProgram {
			cmds = append(cmds, tea.Sequence(finalPause(), tea.Quit))
		}
		m.progress = progressModel.(progress.Model)
		m.currentLog = msg.Message
		return m, tea.Batch(cmds...)

	default:
		return m, nil
	}
}

func (m progressModel) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error() + "\n"
	}

	pad := strings.Repeat(" ", padding)

	viewParts := []string{}

	title := "\n" + pad + m.taskTitle
	viewParts = append(viewParts, title)

	progressBar := pad + m.progress.ViewAs(m.progress.Percent())
	viewParts = append(viewParts, progressBar)

	log := pad + m.currentLog
	if log != "" {
		viewParts = append(viewParts, log)
	}

	help := pad + helpStyle("Press any key to cancel and exit")
	viewParts = append(viewParts, help)

	return strings.Join(viewParts, "\n\n")
}
