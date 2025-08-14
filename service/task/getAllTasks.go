package task_svc

func (s *TaskService) GetAllTasks(status string) ([]*TaskRead, error) {
	tasks := s.taskStorage.GetTasks(status)
	return tasks, nil
}
