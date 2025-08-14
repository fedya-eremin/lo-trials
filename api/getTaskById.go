package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *Server) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 0)
	if err != nil {
		slog.Error("bad request: %v", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task, err := s.taskService.GetTaskById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, err := json.Marshal(task)
	if err != nil {
		slog.Error("marshl failed: %v", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
