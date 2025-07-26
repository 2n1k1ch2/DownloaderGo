package handlers

import (
	"DownloaderGo/internal/tasks"
	"DownloaderGo/pkg"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func CreateTask(w http.ResponseWriter, r *http.Request, tm *tasks.TaskManager) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	task := tm.CreateTask()

	resp := pkg.JSONCreateTaskResponse{TaskID: task.Id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func GetTask(w http.ResponseWriter, r *http.Request, tm *tasks.TaskManager) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tsk, err := tm.FindTask(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(pkg.JSONErrorResponse{Error: err.Error()})
		return
	}

	resp := pkg.JSONTaskStatusResponse{
		Status:     tsk.TaskStatus.String(),
		ArchiveURL: tsk.ArchiveURL,
		Files:      tsk.Links,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func AddLink(w http.ResponseWriter, r *http.Request, tm *tasks.TaskManager) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	url := r.URL.Query().Get("url")
	if id == "" || url == "" {
		http.Error(w, "id and url are required", http.StatusBadRequest)
		return
	}
	tsk, err := tm.AddLink(id, url)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(pkg.JSONErrorResponse{Error: err.Error()})
		return
	}
	resp := pkg.JSONAddLinkResponse{Status: tsk.TaskStatus.String(), Error: ""}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func DownloadArchive(w http.ResponseWriter, r *http.Request, tm *tasks.TaskManager) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	// путь до архива
	archivePath := filepath.Join("files", id, "archive.zip")

	// проверка существования
	if _, err := os.Stat(archivePath); os.IsNotExist(err) {
		http.Error(w, "archive not found", http.StatusNotFound)
		return
	}

	// установка заголовков и отдача файла
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", id))
	http.ServeFile(w, r, archivePath)
}
