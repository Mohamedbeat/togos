package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mohamedbeat/togo/models"
)

var modalStyle = lipgloss.NewStyle().
	Padding(1, 4).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("205")).
	Align(lipgloss.Center).
	Width(50)

type editModel struct {
	inputs  []textInput
	focused int
	index   int // index of the togo being edited
}

// --- Centering support ---
// Add width and height to the root model (main.go must be updated accordingly)
type windowSizer interface {
	GetWindowSize() (int, int)
}

func (e editModel) View(store *models.Store, width, height int) string {
	var sb strings.Builder
	sb.WriteString("Edit Task (press esc to cancel)\n\n")

	for i, input := range e.inputs {
		label := ""
		switch i {
		case 0:
			label = "Title: "
		case 1:
			label = "Description: "
		}
		if i == e.focused {
			label = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(label)
		}
		sb.WriteString(fmt.Sprintf("%s%s\n", label, input.value))
	}

	sb.WriteString("\nPress Enter to save changes | Esc to go back")

	modal := modalStyle.Render(sb.String())

	// Use lipgloss.Place to center the modal
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, modal)
}

func (m model) updateEdit(msg tea.Msg, store *models.Store) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var done bool
	m.edit, cmd, done = m.edit.Update(msg, store)
	if done {
		task := &store.Tasks[m.edit.index]
		store.UpdateTask(task.ID.String(), m.edit.inputs[0].value, m.edit.inputs[1].value)
		m.currentView = "list"
	}
	return m, cmd
}

func (e editModel) Update(msg tea.Msg, store *models.Store) (editModel, tea.Cmd, bool) {
	done := false
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			if s == "enter" && e.focused == len(e.inputs)-1 {
				done = true
				return e, nil, done
			}
			if s == "up" || s == "shift+tab" {
				e.focused--
			} else {
				e.focused++
			}
			if e.focused >= len(e.inputs) {
				e.focused = 0
			} else if e.focused < 0 {
				e.focused = len(e.inputs) - 1
			}
			return e, nil, false
		}
	}
	// Update focused field
	cmd := e.updateInputs(msg)
	return e, cmd, false
}

func (e *editModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(e.inputs))
	for i := range e.inputs {
		if i == e.focused {
			var cmd tea.Cmd
			e.inputs[i], cmd = e.updateInput(msg, e.inputs[i])
			cmds[i] = cmd
		}
	}
	return tea.Batch(cmds...)
}

func (e editModel) updateInput(msg tea.Msg, input textInput) (textInput, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "backspace":
			if len(input.value) > 0 {
				input.value = input.value[:len(input.value)-1]
			}
		default:
			input.value += msg.String()
		}
	}
	return input, nil
}
