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
	if len(t.Content) > 50 {
		return zone.Mark(t.Code+t.Content[0:50], t.Content[0:50])
	} else {
		return zone.Mark(t.Code+t.Content, t.Content)
	}
}
