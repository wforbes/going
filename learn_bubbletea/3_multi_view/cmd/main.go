package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type screenOneModel struct{}

func (m screenOneModel) Init() tea.Cmd {
	return nil
}

func screenOne() screenOneModel {
	return screenOneModel{}
}

func (m screenOneModel) View() string {
	str := "This is the first screen. Press any key to switch to the second screen."
	return str
}

func (m screenOneModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			// any other key switches the screen
			m2 := screenTwo()
			return m2, m2.Init()
		}
	default:
		return m, nil
	}
}

type screenTwoModel struct {
	spinner spinner.Model
	err     error
}

func screenTwo() screenTwoModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Bold(true)
	return screenTwoModel{
		spinner: s,
	}
}

func (m screenTwoModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m screenTwoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			// any other key switches the screen
			m1 := screenOne()
			return m1, m1.Init()
		}
	case error:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m screenTwoModel) View() string {
	str := fmt.Sprintf("\n   %s This is screen two...\n\n", m.spinner.View())
	return str
}

func main() {
	p := tea.NewProgram(screenOne(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
}
