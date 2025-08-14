package api

import task_svc "github.com/fedya-eremin/lo-trials/service/task"

type Server struct {
	taskService *task_svc.TaskService
}

func NewServer(svc *task_svc.TaskService) *Server {
	return &Server{
		taskService: svc,
	}
}
