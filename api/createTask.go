package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	task_svc "github.com/fedya-eremin/lo-trials/service/task"
)

const MB = 1024 * 1024 // 1MB in B

func (s *Server) AddTask(w http.ResponseWriter, r *http.Request) {
	slog.Info("qweqwe")
	r.Body = http.MaxBytesReader(w, r.Body, MB)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var task task_svc.TaskCreate
	err := dec.Decode(&task)
	if err != nil {
		slog.Error("bad request: %v", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newTask, err := s.taskService.AddTask(&task)
	if err != nil {
		slog.Error("internal error: %v", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(newTask)
	if err != nil {
		slog.Error("failed to marshal: %v", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
