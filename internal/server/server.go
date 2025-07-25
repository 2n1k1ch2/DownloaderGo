package server

import (
	"DownloaderGo/internal/tasks"
	"net/http"
)

type Server struct {
	Mux         http.ServeMux
	TaskManager tasks.TaskManager
}

func NewServer() *Server {
	return &Server{}
}
