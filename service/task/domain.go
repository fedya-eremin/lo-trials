package task_svc

import "time"

type TaskCreate struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	AssigneeId  int        `json:"assignee_id"`
	Deadline    *time.Time `json:"deadline"`
	Status      string     `json:"status"`
}

type TaskRead struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	AssigneeId  int        `json:"assignee_id"`
	Deadline    *time.Time `json:"deadline"`
	Status      string     `json:"status"`
	Id          uint64     `json:"id"`
}
