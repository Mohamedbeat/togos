package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mohamedbeat/togo/models"
)

var saveModalStyle = lipgloss.NewStyle().
	Padding(1, 4).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("205")).
	Align(lipgloss.Center).
	Width(40)

type saveModel struct {
}

func (m model) updateSave(msg tea.Msg, store *models.Store) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var done bool
	var save bool
	m.save, cmd, done, save = m.save.Update(msg, store)
	if done {
		if save {
			err := store.SaveStore()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		m.currentView = "list"
		return m, tea.ClearScreen
	}
	return m, cmd
}

func (s saveModel) View(store *models.Store, width, height int) string {
	str := "Save changes ?\n"
	// str += "\ny | n.\n"

	str += "y = Yes, n = No\n"
	modal := saveModalStyle.Render(str)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, modal)
}

func (s saveModel) Update(msg tea.Msg, store *models.Store) (saveModel, tea.Cmd, bool, bool) {
	done := false
	save := false
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y":
			done = true
			save = true
			return s, nil, done, save
		case "n":
			done = true
			save = false
			return s, nil, done, save
		}
		return s, nil, false, false
	}
	return s, nil, false, false
}
