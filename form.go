package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mohamedbeat/togo/models"
)

var formModalStyle = lipgloss.NewStyle().
	Padding(1, 4).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("205")).
	Align(lipgloss.Center).
	Width(50)

type formModel struct {
	inputs  []textInput
	focused int
}

// Text input field
type textInput struct {
	value string
}

// Delegate form view updates
func (m model) updateForm(msg tea.Msg, store *models.Store) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var done bool
	m.form, cmd, done = m.form.Update(msg, store)
	if done {
		m.currentView = "list"
	}
	return m, cmd
}

func (f formModel) View(store *models.Store, width, height int) string {
	var sb strings.Builder
	sb.WriteString("New Task Form (press esc to cancel)\n\n")

	for i, input := range f.inputs {
		label := ""
		switch i {
		case 0:
			label = "Title: "
		case 1:
			label = "Description: "
		}

		// Highlight focused field
		if i == f.focused {
			label = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(label)
		}

		sb.WriteString(fmt.Sprintf("%s%s\n", label, input.value))
	}

	sb.WriteString("\nPress enter to submit")

	modal := formModalStyle.Render(sb.String())
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, modal)
}

func (f formModel) Update(msg tea.Msg, store *models.Store) (formModel, tea.Cmd, bool) {
	done := false
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab", "up", "down":
			// Handle field navigation
			if msg.String() == "up" || msg.String() == "shift+tab" {
				f.focused--
			} else {
				f.focused++
			}
			if f.focused >= len(f.inputs) {
				f.focused = 0
			} else if f.focused < 0 {
				f.focused = len(f.inputs) - 1
			}
			return f, nil, false
		case "enter":
			if f.focused == len(f.inputs)-1 {
				store.NewTask(models.NewTaskDto{
					Title: f.inputs[0].value,
					Desc:  f.inputs[1].value,
				})
				for i := 0; i < len(f.inputs); i++ {
					f.inputs[i].value = ""
				}
				done = true
			}
			return f, nil, done
		}
	}

	// Update focused field
	cmd := f.updateInputs(msg)
	return f, cmd, false
}

func (m *formModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only update the focused field
	for i := range m.inputs {
		if i == m.focused {
			var cmd tea.Cmd
			m.inputs[i], cmd = m.updateInput(msg, m.inputs[i])
			cmds[i] = cmd
		}
	}

	return tea.Batch(cmds...)
}

func (m formModel) updateInput(msg tea.Msg, input textInput) (textInput, tea.Cmd) {
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
