package cluster

import (
	"github.com/l-angel/tunnel/worker"
)

type Cluster struct {
	electionSignal chan bool
	w              *worker.Worker
}

func NewCluster(signal chan bool, w *worker.Worker) *Cluster {
	return &Cluster{electionSignal: signal, w: w}
}

func (c *Cluster) Boot() {
	c.watchElect()
}

func (c *Cluster) watchElect() {
	go func() {
		for {
			select {
			case selected := <-c.electionSignal:
				if selected {
					c.w.StartUp()
				} else {
					c.w.Stop()
				}
			}
		}
	}()
}
