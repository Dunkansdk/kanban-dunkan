package task

import zone "github.com/lrstanley/bubblezone"

type TaskStatus struct {
	ID   int
	Name string
}

type Task struct {
	ID      int64
	Code    string
	Status  TaskStatus
	Name    string
	Content string
}

// Method mandatory in bubbletea.
func (t Task) FilterValue() string {
	return zone.Mark(t.Code+t.Name, t.Name)
}

func (t Task) Title() string {
	return zone.Mark(t.Code+t.Name, t.Name)
}

func (t Task) Description() string {
	return zone.Mark(t.Code, t.Code)
}
