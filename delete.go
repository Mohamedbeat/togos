package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mohamedbeat/togo/models"
)

var deleteModalStyle = lipgloss.NewStyle().
	Padding(1, 4).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("1")).
	Align(lipgloss.Center).
	Width(50)

type deleteModel struct {
	index int
	id    string
	title string
}

func (d deleteModel) View(store *models.Store, width, height int) string {
	str := "Delete this task?\n\n"
	str += lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true).Render(d.title) + "\n\n"
	str += "y = Yes, n = No"
	modal := deleteModalStyle.Render(str)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, modal)
}

func (d deleteModel) Update(msg tea.Msg, store *models.Store) (deleteModel, tea.Cmd, bool, bool) {
	done := false
	deleteConfirmed := false
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y":
			err := store.DeleteTask(d.id)
			if err != nil {
				fmt.Println(err)
			}
			done = true
			deleteConfirmed = true
			return d, nil, done, deleteConfirmed
		case "n":
			done = true
			deleteConfirmed = false
			return d, nil, done, deleteConfirmed
		}
	}
	return d, nil, false, false
}

func (m model) updateDelete(msg tea.Msg, store *models.Store) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var done bool
	m.delete, cmd, done, _ = m.delete.Update(msg, store)
	if done {
		m.currentView = "list"
		return m, tea.ClearScreen
	}
	return m, cmd
}
