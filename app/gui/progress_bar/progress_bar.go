package progressbar

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type progressWriter struct {
	onProgress func(ProgressMsg)
	program    *tea.Program
}

func (p *progressWriter) Run(task func(inputChannel chan ProgressMsg) error) {
	inputChannel := make(chan ProgressMsg)

	go task(inputChannel)

	go func() {
		for msg := range inputChannel {
			// log.Print(msg)
			p.program.Send(msg)
		}
	}()

	if _, err := p.program.Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}

func NewProgressWriter(taskTitle string) *progressWriter {
	m := progressModel{
		progress:  progress.New(progress.WithDefaultGradient()),
		taskTitle: taskTitle,
	}

	return &progressWriter{
		program: tea.NewProgram(m),
	}
}
