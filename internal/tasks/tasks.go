package tasks

import (
	"DownloaderGo/internal/archiver"
	"DownloaderGo/internal/fetcher"
	"fmt"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"sync"
)

type Status int

const (
	Pending Status = iota
	Active
	Rejected
	Completed
)

func (s Status) String() string {
	switch s {
	case Pending:
		return "pending"
	case Active:
		return "active"
	case Rejected:
		return "rejected"
	case Completed:
		return "completed"
	default:
		return "unknown"
	}
}

type Task struct {
	sync.Mutex
	Id         string
	TaskStatus Status
	Links      []File
	ArchiveURL string
}
type File struct {
	URL    string `json:"url"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func NewTask() *Task {
	task := new(Task)

	task.Id = uuid.New().String()
	task.TaskStatus = Pending
	task.Links = []File{}
	task.ArchiveURL = ""
	return task
}

func (tsk *Task) Run() {
	taskDir := filepath.Join("files", tsk.Id)
	err := os.MkdirAll(taskDir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	tsk.Lock()
	tsk.TaskStatus = Active
	tsk.Unlock()
	var downloadedFiles []string

	for i := range tsk.Links {
		url := tsk.Links[i].URL
		if !fetcher.IsAllowedFile(url) {
			tsk.Links[i].Status = Rejected.String()
			continue
		}

		data, err := fetcher.DownloadFile(url)
		if err != nil {
			tsk.Links[i].Status = Rejected.String()
			continue
		}

		filename := fmt.Sprintf("files/%s/%s_%d%s", tsk.Id, tsk.Id, i, filepath.Ext(url))
		err = os.WriteFile(filename, data, 0644)
		if err != nil {
			tsk.Links[i].Status = Rejected.String()
			tsk.Links[i].Error = err.Error()
			continue
		}

		downloadedFiles = append(downloadedFiles, filename)
		tsk.Links[i].Status = Completed.String()
	}

	if len(downloadedFiles) > 0 {
		archivePath := filepath.Join("files", tsk.Id, "archive.zip")
		err := archiver.CreateZip(archivePath, downloadedFiles)
		if err == nil {
			tsk.ArchiveURL = fmt.Sprintf("http://localhost:8080/Download?id=%s", tsk.Id)
		}
	}

	tsk.Lock()
	tsk.TaskStatus = Completed
	tsk.Unlock()
}
