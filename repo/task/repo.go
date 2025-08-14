package task_repo

import (
	"github.com/fedya-eremin/lo-trials/service/task"
	"github.com/fedya-eremin/lo-trials/store"
)

type TaskRepo struct {
	storage *store.InMemoryStore[task_svc.TaskCreate]
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{
		storage: store.NewInMemoryStore[task_svc.TaskCreate](),
	}
}
