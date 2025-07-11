package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mohamedbeat/togo/models"
)

type listModel struct {
	cursor int // which to-do list item our cursor is pointing at
	table  table.Model
}

func (m model) updateList(msg tea.Msg, store *models.Store) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg, store)
	return m, cmd
}

func (l listModel) View(store *models.Store, width, height int) string {
	// Prepare table rows from store.Tasks
	rows := []table.Row{}
	for _, t := range store.Tasks {
		status := " "
		if t.Done {
			status = "x"
		}
		createdAt := t.CreatedAt.Format("2006-01-02 15:04")
		doneAt := ""
		if t.DoneAt != nil {
			doneAt = t.DoneAt.Format("2006-01-02 15:04")
		}
		rows = append(rows, table.Row{status, t.Title, t.Desc, createdAt, doneAt})
	}

	columns := []table.Column{
		{Title: "Status", Width: 6},
		{Title: "Title", Width: 30},
		{Title: "Description", Width: 30},
		{Title: "CreatedAt", Width: 17},
		{Title: "DoneAt", Width: 17},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
	t.SetHeight(height - 6) // leave space for header/footer
	t.SetWidth(width - 8)
	t.SetCursor(l.cursor)

	header := "Hello\n\n"
	footer := "\nPress q to quit  Press a to add  Press w|crtl+w to save  Press d to delete  Press e to edit.\n"

	return fmt.Sprintf("%s%s%s", header, t.View(), footer)
}

func (l listModel) Update(msg tea.Msg, store *models.Store) (listModel, tea.Cmd) {
	total := len(store.Tasks)
	t := l.table
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return l, tea.Quit
		case "up", "k":
			if l.cursor > 0 {
				l.cursor--
			}
		case "down", "j":
			if l.cursor < total-1 {
				l.cursor++
			}
		case "enter", " ":
			if l.cursor >= 0 && l.cursor < total {
				store.Tasks[l.cursor].ToggleDone()
			}
		}
	}

	// Update the table's cursor
	t.SetCursor(l.cursor)
	l.table = t

	return l, cmd
}
