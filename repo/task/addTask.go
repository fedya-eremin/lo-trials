package task_repo

import (
	task_svc "github.com/fedya-eremin/lo-trials/service/task"
)

func (r *TaskRepo) AddTask(task *task_svc.TaskCreate) *task_svc.TaskRead {
	id := r.storage.Add(*task)
	return &task_svc.TaskRead{
		Id:          id,
		Name:        task.Name,
		Description: task.Description,
		AssigneeId:  task.AssigneeId,
		Status:      task.Status,
		Deadline:    task.Deadline,
	}
}
