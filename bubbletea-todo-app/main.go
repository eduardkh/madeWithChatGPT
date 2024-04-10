package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	todos  []string // Tasks to do
	done   []bool   // Completion status of tasks
	cursor int      // Current cursor position
	input  string   // User input for new tasks
}

func (m model) Init() tea.Cmd {
	// Initialize without any commands
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.input != "" {
				m.todos = append(m.todos, m.input)
				m.done = append(m.done, false)
				m.input = "" // Reset input after adding a task
			}
		case tea.KeyUp, tea.KeyCtrlK:
			if m.cursor > 0 {
				m.cursor-- // Move the cursor up
			}
		case tea.KeyDown, tea.KeyCtrlJ:
			if m.cursor < len(m.todos)-1 {
				m.cursor++ // Move the cursor down
			}
		case tea.KeyCtrlD:
			if len(m.todos) > 0 && m.cursor < len(m.todos) {
				m.done[m.cursor] = !m.done[m.cursor] // Toggle task completion
			}
		case tea.KeyBackspace:
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1] // Remove the last character
			}
		case tea.KeySpace:
			m.input += " " // Explicitly handle space addition
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit // Quit the application
		default:
			if msg.Type == tea.KeyRunes {
				m.input += msg.String() // Append other character inputs
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "What do you need to do today?\n\n"
	for i, todo := range m.todos {
		cursor := " "
		if m.cursor == i {
			cursor = "->"
		}
		done := " "
		if m.done[i] {
			done = "X"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, done, todo)
	}
	s += "\n>>: " + m.input + ""
	return s
}

func main() {
	initialModel := model{}
	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
