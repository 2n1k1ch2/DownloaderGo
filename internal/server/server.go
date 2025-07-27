package server

import (
	"DownloaderGo/internal/handlers"
	"DownloaderGo/internal/tasks"
	"net/http"
)

type Server struct {
	http.Handler
	Mux         http.ServeMux
	TaskManager *tasks.TaskManager
}

func NewServer() *Server {

	mux := http.NewServeMux()
	s := &Server{
		Mux:         *mux,
		Handler:     mux,
		TaskManager: tasks.NewTaskManager(),
	}

	// Регистрация ручек
	mux.HandleFunc("/create-task", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTask(w, r, s.TaskManager)
	})
	mux.HandleFunc("/get-task", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTask(w, r, s.TaskManager)
	})
	mux.HandleFunc("/add-link", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddLink(w, r, s.TaskManager)
	})
	mux.HandleFunc("/Download", func(w http.ResponseWriter, r *http.Request) {
		handlers.DownloadArchive(w, r, s.TaskManager)
	})

	return s
}
