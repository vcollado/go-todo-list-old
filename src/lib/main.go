package lib

type Task struct {
	Name        string
	Description string
	Done        bool
}

// GetName returns the name of the task
func (t Task) GetName() string {
	return t.Name
}

// IsDone check if the tasks is done
func (t Task) IsDone() bool {
	return t.Done
}
