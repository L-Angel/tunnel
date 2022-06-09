package worker

import (
	"github.com/l-angel/tunnel/cfg"
	"github.com/l-angel/tunnel/registry"
)

type Worker struct {
	tasks map[string]*Task
	r     registry.Registry
}

func NewWorkerWithCfg(r registry.Registry) *Worker {
	tasks := make(map[string]*Task)
	for _, t := range cfg.C.Tasks {
		tasks[t.Id] = NewTask(t.Id, t, r)
	}
	return &Worker{
		tasks: tasks,
	}

}
func NewEmptyWorker() *Worker {
	return &Worker{tasks: make(map[string]*Task)}
}

func (self *Worker) Join(id string, t *Task) {
	if t != nil {
		self.tasks[id] = t
	}
}

func (self *Worker) StartUp() {
	for _, t := range self.tasks {
		_ = t.StartUp()
	}
}

/*
校验任务状态 Task.Status 和 Task 之间实际的工作状态校验
*/
func (self *Worker) Refresh() {
	for _, s := range self.tasks {
		if s.Status == Waiting {
			_ = s.StartUp()
		}
	}
}

func (self *Worker) Stop() {
	for _, t := range self.tasks {
		t.Stop()
	}
}
