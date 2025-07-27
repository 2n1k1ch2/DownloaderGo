package handlers

import (
	"DownloaderGo/internal/tasks"
	"DownloaderGo/pkg"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CreateTask(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/create-task", nil)
	tm := tasks.NewTaskManager()
	CreateTask(w, r, tm)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 200 OK, got %d", w.Code)
	}

	var resp pkg.JSONCreateTaskResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	if resp.TaskID == "" {
		t.Fatalf("task ID is empty")
	}
}
func Test_CreateTaskWrongMethod(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/create-task", nil)
	tm := tasks.NewTaskManager()
	CreateTask(w, r, tm)
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", w.Code)
	}
}

func Test_GetTask(t *testing.T) {
	w := httptest.NewRecorder()
	tm := tasks.NewTaskManager()
	task := tm.CreateTask()
	r := httptest.NewRequest(http.MethodGet, "/get-task?id="+task.Id, nil)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	GetTask(w, r, tm)
	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}
func Test_GetTaskWrongMethod(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/get-task", nil)
	tm := tasks.NewTaskManager()
	GetTask(w, r, tm)
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", w.Code)
	}
}
func Test_GetTaskEmptyID(t *testing.T) {
	w := httptest.NewRecorder()
	tm := tasks.NewTaskManager()
	r := httptest.NewRequest(http.MethodGet, "/get-task?id=", nil)
	GetTask(w, r, tm)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}
}
func Test_GetTaskNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	tm := tasks.NewTaskManager()
	r := httptest.NewRequest(http.MethodGet, "/get-task?id=1", nil)
	GetTask(w, r, tm)
	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", w.Code)
	}
}

func Test_AddLink(t *testing.T) {
	w := httptest.NewRecorder()
	tm := tasks.NewTaskManager()
	task := tm.CreateTask()
	r := httptest.NewRequest(http.MethodPost, "/add-link?id="+task.Id+"&url="+
		"https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf", nil)
	AddLink(w, r, tm)
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}
func Test_AddLinkWrongMethod(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/add-link", nil)
	tm := tasks.NewTaskManager()
	AddLink(w, r, tm)
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", w.Code)
	}
}
func Test_AddLinkEmptyID(t *testing.T) {
	w := httptest.NewRecorder()
	tm := tasks.NewTaskManager()
	r := httptest.NewRequest(http.MethodPost, "/add-link?id=&url=", nil)
	AddLink(w, r, tm)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)

	}
}
func Test_AddLinkNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	tm := tasks.NewTaskManager()
	r := httptest.NewRequest(http.MethodPost, "/add-link?id=1&url=https://www.w3.org/WAI"+
		"/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf", nil)
	AddLink(w, r, tm)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 404, got %d", w.Code)
	}
}
