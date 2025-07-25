package pkg

import "DownloaderGo/internal/tasks"

// Ответ для POST /create-task
type JSONCreateTaskResponse struct {
	TaskID string `json:"task_id"`
}

// Тело запроса для POST /add-link
type JSONAddLinkRequest struct {
	TaskID string `json:"task_id"`
	URL    string `json:"url"`
}

// Ответ для POST /add-link
type JSONAddLinkResponse struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

// Ответ для GET /task-status
type JSONTaskStatusResponse struct {
	Status     string       `json:"status"`                // Общий статус задачи
	ArchiveURL string       `json:"archive_url,omitempty"` // Ссылка на архив, если готов
	Files      []tasks.File `json:"files,omitempty"`       // Статусы по каждому файлу
}

// Ответ при ошибках (общий)
type JSONErrorResponse struct {
	Error string `json:"error"`
}
