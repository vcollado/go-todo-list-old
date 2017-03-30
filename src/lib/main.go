package lib

type Task struct {
	Name        string
	Description string
	Done        bool
}

func (t Task) GetName() string {
	return t.Name
}

func (t Task) IsDone() bool {
	return t.Done
}
