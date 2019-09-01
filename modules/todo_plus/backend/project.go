package backend

type Task struct {
	ID        int
	Completed bool
	Content   string
}

type Project struct {
	ID   int
	Name string

	Index   int
	Tasks   []Task
	Err     error
	backend Backend
}

func (proj *Project) IsLast() bool {
	return proj.Index >= len(proj.Tasks)-1
}

func NewProject(id int, backend Backend) *Project {
	proj := backend.NewProject(id)
	proj.backend = backend
	return proj
}

func (proj *Project) loadTasks() {
	Tasks, err := proj.backend.LoadTasks(proj.ID)
	proj.Err = err
	proj.Tasks = Tasks
}

func (proj *Project) LongestLine() int {
	maxLen := 0

	for _, task := range proj.Tasks {
		if len(task.Content) > maxLen {
			maxLen = len(task.Content)
		}
	}

	return maxLen
}

func (proj *Project) currentTask() *Task {
	if proj.Index < 0 {
		return nil
	}

	return &proj.Tasks[proj.Index]
}

func (proj *Project) CloseSelectedTask() {
	currTask := proj.currentTask()

	if currTask != nil {
		proj.backend.CloseTask(currTask)
		proj.loadTasks()
	}
}

func (proj *Project) DeleteSelectedTask() {
	currTask := proj.currentTask()

	if currTask != nil {
		proj.backend.DeleteTask(currTask)

		proj.loadTasks()
	}
}
