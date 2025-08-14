package task_repo

import (
	"fmt"

	task_svc "github.com/fedya-eremin/lo-trials/service/task"
)

func (r *TaskRepo) GetTaskById(id uint64) (*task_svc.TaskRead, error) {
	task, ok := r.storage.Get(id)
	if !ok {
		return nil, fmt.Errorf("Task not found")
	}
	return &task_svc.TaskRead{
		Id:          id,
		Name:        task.Name,
		Description: task.Description,
		AssigneeId:  task.AssigneeId,
		Status:      task.Status,
		Deadline:    task.Deadline,
	}, nil
}
