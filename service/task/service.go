package task_svc

type TaskStorage interface {
	GetTasks(status string) []*TaskRead
	AddTask(task *TaskCreate) *TaskRead
	GetTaskById(id uint64) (*TaskRead, error)
}

type TaskService struct {
	taskStorage TaskStorage
}

func NewTaskService(ts TaskStorage) *TaskService {
	return &TaskService{
		taskStorage: ts,
	}
}
