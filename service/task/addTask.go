package task_svc

func (s *TaskService) AddTask(t *TaskCreate) (*TaskRead, error) {
	task := s.taskStorage.AddTask(t)
	return task, nil
}
