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

	// Executing the main task
	go task(inputChannel)

	// Sending the program an information to update the progress bar
	go func() {
		for msg := range inputChannel {
			p.program.Send(msg)
		}
	}()

	if _, err := p.program.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func NewProgressWriter(taskTitle string) *progressWriter {
	m := progressModel{
		progress:  progress.New(progress.WithGradient("#009245", "#FCEE21")),
		taskTitle: taskTitle,
	}

	return &progressWriter{
		program: tea.NewProgram(m),
	}
}
