package worker

import (
	"github.com/l-angel/tunnel/cfg"
	"github.com/l-angel/tunnel/log"
	"github.com/l-angel/tunnel/registry"
)

type Worker struct {
	tasks map[string]*Task
	r     registry.Registry
}

func NewWorkerWithCfg(r registry.Registry) *Worker {
	tasks := make(map[string]*Task)
	for _, t := range cfg.C.Tasks {
		if !t.Enable {
			continue
		}
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
		err := t.StartUp()
		if err != nil {
			log.Error("task [", t.id, "] start failure.", err)
		}
	}
}

/*
校验任务状态 Task.Status 和 Task 之间实际的工作状态校验
*/
func (self *Worker) Refresh() {
	for _, t := range self.tasks {
		if t.Status == Waiting {
			err := t.StartUp()
			if err != nil {
				log.Error("task [", t.id, "] start failure.", err)
			}
		}
	}
}

func (self *Worker) Stop() {
	for _, t := range self.tasks {
		t.Stop()
	}
}
