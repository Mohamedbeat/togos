package models

import (
	"encoding/json"
	"fmt"

	"github.com/mohamedbeat/togo/utils"
)

type Store struct {
	Tasks []Task `json:"tasks"`
}

func NewStore() (*Store, error) {
	bts, err := utils.ReadFile()
	if err != nil {
		return nil, err
	}

	var store Store

	// Unmarshal the JSON data into the store
	err = json.Unmarshal(bts, &store)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &store, nil

}

func (s *Store) NewTask(t NewTaskDto) {

	newTask := NewTask(t)
	s.Tasks = append(s.Tasks, *newTask)
}

func (s *Store) FindTask(id string) (*Task, int, error) {
	for i := range s.Tasks {
		if s.Tasks[i].ID.String() == id {
			return &s.Tasks[i], i, nil
		}
	}

	return nil, 0, fmt.Errorf("Task with id %s not found.", id)
}

func (s *Store) UpdateTask(id string, title, desc string) (*Task, error) {
	t, _, err := s.FindTask(id)

	if err != nil {
		return nil, err
	}
	fmt.Println(t)
	if len(title) != 0 {
		t.Title = title
	}
	if len(desc) != 0 {
		t.Desc = desc
	}

	fmt.Println(t)
	return t, err

}

func (s *Store) ToggleDoneTask(id string) error {
	t, _, err := s.FindTask(id)

	if err != nil {
		return err
	}

	_ = t.ToggleDone()
	return nil
}

func (s *Store) DeleteTask(id string) error {
	_, i, err := s.FindTask(id)

	if err != nil {
		return err
	}
	s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
	return nil

}
func (s *Store) DeleteAllTasks() {
	s.Tasks = []Task{}
}

func (s *Store) SaveStore() error {

	// Marshal the store to JSON with indentation
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling to JSON: %w", err)
	}

	return utils.WriteFile(string(data))

}
