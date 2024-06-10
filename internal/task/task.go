package task

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
	return t.Name
}

func (t Task) Title() string {
	return t.Name
}

func (t Task) Description() string {
	if len(t.Content) > 50 {
		return t.Content[0:50]
	} else {
		return t.Content
	}
}
