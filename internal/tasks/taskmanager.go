package tasks

import (
	"errors"
	"sync"
)

type TaskManager struct {
	activeTasks chan struct{} // семафор на 3 задачи
	tasks       map[string]*Task
	sync.RWMutex
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		activeTasks: make(chan struct{}, 3),
		tasks:       make(map[string]*Task),
	}
}

func (tm *TaskManager) TryStartTask(task *Task) error {
	select {
	case tm.activeTasks <- struct{}{}:
		go tm.runTask(task)
		return nil
	default:
		return errors.New("server busy: max active tasks reached")
	}
}

func (tm *TaskManager) runTask(task *Task) {
	defer func() { <-tm.activeTasks }()
	task.Run()
}

func (tm *TaskManager) CreateTask() (tsk *Task) {
	tsk = NewTask()
	tm.RWMutex.Lock()
	defer tm.RWMutex.Unlock()
	tm.tasks[tsk.Id] = tsk
	return
}
func (tm *TaskManager) FindTask(id string) (*Task, error) {
	tm.RWMutex.RLock()
	defer tm.RWMutex.RUnlock()
	if tsk, ok := tm.tasks[id]; ok {
		return tsk, nil
	}
	return nil, errors.New("task not found")
}
func (tm *TaskManager) AddLink(id string, url string) (*Task, error) {
	tm.RWMutex.RLock()
	defer tm.RWMutex.RUnlock()
	tsk, ok := tm.tasks[id]
	if !ok {
		return nil, errors.New("task does not exist")
	}

	tsk.Lock()
	defer tsk.Unlock()

	if len(tsk.Links) >= 3 {
		return nil, errors.New("maximum 3 links allowed")
	}

	tsk.Links = append(tsk.Links, File{URL: url, Status: Pending.String()})

	if len(tsk.Links) == 3 {
		err := tm.TryStartTask(tsk)
		if err != nil {
			return nil, err
		}
	}
	return tsk, nil
}
