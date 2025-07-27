package tasks

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateTask(t *testing.T) {
	tm := NewTaskManager()
	task := tm.CreateTask()
	if task == nil {
		t.Error("task creation failed")
	}

}
func TestFindTask(t *testing.T) {
	tm := NewTaskManager()
	task := tm.CreateTask()

	if _, err := tm.FindTask(task.Id); err != nil {
		t.Error("task does not exists")
	}
}
func TestFindTaskWrongId(t *testing.T) {
	tm := NewTaskManager()
	if _, err := tm.FindTask("1"); err == nil {
		t.Error("task should not exists")
	}
}
func TestAddLink(t *testing.T) {
	tm := NewTaskManager()
	task := tm.CreateTask()
	tm.AddLink(task.Id, "1")
	if len(task.Links) != 1 {
		t.Error("task links do not match")
	}
}
func TestAddLinkMoreThan3(t *testing.T) {
	tm := NewTaskManager()
	task := tm.CreateTask()
	for i := 0; i < 3; i++ {
		tm.AddLink(task.Id, fmt.Sprintf("%d", i))
	}
	_, err := tm.AddLink(task.Id, fmt.Sprintf("%d", 3))
	if err == nil {
		t.Error("task links do not match")
	}
}
func TestAddlinkTaskNotExist(t *testing.T) {
	tm := NewTaskManager()
	_, err := tm.AddLink("1", fmt.Sprintf("%d", 3))
	if err == nil {
		t.Error("task exists , but should not")
	}
}

func TestTask_Run(t *testing.T) {
	tm := NewTaskManager()
	tsk := tm.CreateTask()
	_, err := tm.AddLink(tsk.Id, "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf")
	if err != nil {
		t.Error(err)
	}
	_, err = tm.AddLink(tsk.Id, "1")
	if err != nil {
		t.Error(err)
	}
	tsk.Run()
	if err != nil {
		t.Error(err)
	}
	if tsk.TaskStatus != Completed {
		t.Errorf("expected task status Completed, got %v", tsk.TaskStatus)
	}

	if tsk.ArchiveURL == "" {
		t.Error("ArchiveURL not set")
	}

	if tsk.Links[0].Status != Completed.String() {
		t.Errorf("expected file status Completed, got %s", tsk.Links[0].Status)
	}

	archivePath := filepath.Join("files", tsk.Id, "archive.zip")
	if _, err := os.Stat(archivePath); err != nil {
		t.Errorf("archive not created: %v", err)
	}

	// Чистим
	err = os.RemoveAll("files")
	if err != nil {
		t.Error(err)
	}
}
