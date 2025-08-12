package task

type TaskServers interface {
	CreateTask(task *Task) (*Task, error)
	GetAllTask() ([]Task, error)
	GetTaskByID(id int) (Task, error)
	UpdataTask(id int, expression string) (Task, error)
	DeleteTask(id int) error
	GetTasksForUser(userID int) ([]Task, error)
}

type taskService struct {
	repo TaskRepository
}

// GetTasksForUser implements TaskServers.
func (s *taskService) GetTasksForUser(userID int) ([]Task, error) {
	return s.repo.GetTasksByUserID(userID)
}

func NewTaskService(r TaskRepository) TaskServers {
	return &taskService{repo: r}
}
func (s *taskService) CreateTask(task *Task) (*Task, error) {
	if err := s.repo.CreateTask(task); err != nil {
		return nil, err
	}
	return task, nil
}
func (s *taskService) GetAllTask() ([]Task, error) {
	return s.repo.GetAllTask()

}

func (s *taskService) GetTaskByID(id int) (Task, error) {
	return s.repo.GetTaskByID(id)

}

func (s *taskService) UpdataTask(id int, expression string) (Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	task.Task = expression

	if err := s.repo.UpdateTask(task); err != nil {
		return Task{}, err

	}
	return task, nil

}

func (s *taskService) DeleteTask(id int) error {
	return s.repo.DeleteTask(id)
}
