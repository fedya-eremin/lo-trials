package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (s *Server) GetTasksByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	tasks, err := s.taskService.GetAllTasks(status)
	if err != nil {
		slog.Error("failed to get tasks: %v", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(tasks)
	if err != nil {
		slog.Error("marshl failed: %v", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
