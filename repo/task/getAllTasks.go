package task_repo

import task_svc "github.com/fedya-eremin/lo-trials/service/task"

func (r *TaskRepo) GetTasks(status string) []*task_svc.TaskRead {
	tasks := r.storage.Filter(func(id uint64, value task_svc.TaskCreate) bool {
		return status == "" || value.Status == status
	})
	result := make([]*task_svc.TaskRead, 0)
	for id, task := range tasks {
		result = append(result, &task_svc.TaskRead{
			Id:          id,
			Name:        task.Name,
			Description: task.Description,
			AssigneeId:  task.AssigneeId,
			Status:      task.Status,
			Deadline:    task.Deadline,
		})
	}
	return result
}
