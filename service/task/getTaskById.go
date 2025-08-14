package task_svc

func (s *TaskService) GetTaskById(id uint64) (*TaskRead, error) {
	task, err := s.taskStorage.GetTaskById(id)
	if err != nil {
		// TODO log
		return nil, err
	}
	return task, nil
}
