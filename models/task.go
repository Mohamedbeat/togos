package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	Desc      string     `json:"desc"`
	Done      bool       `json:"done"`
	CreatedAt time.Time  `json:"createdAt"`
	DoneAt    *time.Time `json:"doneAt"`
}

type NewTaskDto struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func NewTask(task NewTaskDto) *Task {
	return &Task{ID: uuid.New(), Title: task.Title,
		Desc:      task.Desc,
		Done:      false,
		CreatedAt: time.Now(),
		DoneAt:    nil}
}
func (t *Task) ToggleDone() *Task {
	// var done bool

	t.Done = !t.Done
	if t.Done {
		time := time.Now()
		t.DoneAt = &time
	} else {
		t.DoneAt = nil
	}
	return t
}

func (t *Task) UpdateTitle(title string) *Task {
	t.Title = title
	return t
}

func (t *Task) UpdateDesc(desc string) *Task {
	t.Desc = desc
	return t
}

func (t *Task) String(cursor, checked string) string {
	//  ╭──────────────╮
	//  │  Nerd Box    │
	//  ╰──────────────╯
	var s strings.Builder

	s.WriteString(fmt.Sprintf("%s [%s] %s", cursor, checked, t.Title))

	// if len(t.Desc) != 0 {
	// 	s += fmt.Sprintf(" | %s", t.Desc)
	// }

	year, month, day := t.CreatedAt.Date()
	s.WriteString(fmt.Sprintf(" | Created At: %d-%s-%d", year, month, day))
	if t.DoneAt != nil {
		year, month, day := t.DoneAt.Date()

		s.WriteString(fmt.Sprintf("  Done At: %d-%s-%d", year, month, day))
	}
	s.WriteString(fmt.Sprintf("\n"))
	if len(t.Desc) != 0 {
		s.WriteString("  Description:\n")
		s.WriteString("  ╭────────────────────────────────────────╮\n")
		s.WriteString(fmt.Sprintf("    %s\n", t.Desc))

		s.WriteString("  ╰────────────────────────────────────────╯\n")
	}
	// s.WriteString(fmt.Sprintf("\n"))

	s.WriteString("  ---------------------------------\n")
	// s.WriteString(fmt.Sprintf("\n"))
	return s.String()
}
