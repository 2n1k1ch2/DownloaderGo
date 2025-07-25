package tasks

import (
	"github.com/google/uuid"
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
}

func NewTask() *Task {
	task := new(Task)

	task.Id = uuid.New().String()
	task.TaskStatus = Pending
	task.Links = []File{}
	task.ArchiveURL = ""
	return task
}

func (tsk *Task) run() {

}
