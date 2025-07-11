package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mohamedbeat/togo/models"
)

type model struct {
	currentView string
	store       *models.Store
	list        listModel
	form        formModel
	save        saveModel
	edit        editModel
	delete      deleteModel
	width       int
	height      int
}

func initialModel() (model, error) {
	s, err := models.NewStore()
	if err != nil {
		return model{}, err
	}
	return model{
		currentView: "list",
		store:       s,
		list:        listModel{cursor: 0},
		form: formModel{
			inputs: []textInput{
				textInput{},
				textInput{},
			},
		},
		save:   saveModel{},
		edit:   editModel{},
		delete: deleteModel{},
	}, nil
}
func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return tea.ClearScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {
		case "a":
			if m.currentView == "list" {
				m.currentView = "form"
				return m, nil
			}
		case "w", "crtl+w":
			if m.currentView == "list" {
				m.currentView = "save"
				return m, nil
			}
		case "e":
			if m.currentView == "list" && len(m.store.Tasks) > 0 && m.list.cursor >= 0 && m.list.cursor < len(m.store.Tasks) {
				t := m.store.Tasks[m.list.cursor]
				m.edit = editModel{
					inputs: []textInput{
						textInput{value: t.Title},
						textInput{value: t.Desc},
					},
					focused: 0,
					index:   m.list.cursor,
				}
				m.currentView = "edit"
				return m, nil
			}
		case "d":
			if m.currentView == "list" && len(m.store.Tasks) > 0 && m.list.cursor >= 0 && m.list.cursor < len(m.store.Tasks) {
				t := m.store.Tasks[m.list.cursor]
				m.delete = deleteModel{
					index: m.list.cursor,
					id:    t.ID.String(),
					title: t.Title,
				}
				m.currentView = "delete"
				return m, nil
			}
		case "esc":
			if m.currentView == "form" || m.currentView == "edit" || m.currentView == "delete" {
				m.currentView = "list"
				return m, nil
			}

		}
		switch m.currentView {
		case "list":
			return m.updateList(msg, m.store)
		case "form":
			return m.updateForm(msg, m.store)
		case "save":
			return m.updateSave(msg, m.store)
		case "edit":
			return m.updateEdit(msg, m.store)
		case "delete":
			return m.updateDelete(msg, m.store)
		default:
			return m, nil
		}
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

// Delegate list view updates

func (m model) View() string {
	switch m.currentView {
	case "list":
		return m.list.View(m.store, m.width, m.height)
	case "form":
		return m.form.View(m.store, m.width, m.height)
	case "save":
		return m.save.View(m.store, m.width, m.height)
	case "edit":
		return m.edit.View(m.store, m.width, m.height)
	case "delete":
		return m.delete.View(m.store, m.width, m.height)
	default:
		return "Unknown view"
	}
}

func main() {
	m, err := initialModel()
	if err != nil {
		fmt.Printf("there's been an error: %v", err)
		os.Exit(1)
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("there's been an error: %v", err)
		os.Exit(1)
	}
}
