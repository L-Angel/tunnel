package bootstrap

import (
	_ "github.com/l-angel/tunnel/cfg"
	"github.com/l-angel/tunnel/registry"

	"github.com/l-angel/tunnel/bootstrap/cluster"
	"github.com/l-angel/tunnel/cfg"
	"github.com/l-angel/tunnel/election"
	"github.com/l-angel/tunnel/worker"
	"strings"
)

type Bootstrap struct {
	electionSignal  chan bool
	electionService *election.Election

	// worker
	w *worker.Worker
	// cluster
	c *cluster.Cluster

	r registry.Registry
}

// Boot /** worker must be loaded before cluster election
func (s *Bootstrap) Boot() {
	s.loadCfg()
	s.loadLog()
	s.loadRegistry()
	s.loadWorker()
	s.loadElection()

	_ = s.electionService.Elect()
	s.c.Boot()
}

func (s *Bootstrap) loadCfg() {

}

func (s *Bootstrap) loadLog() {

}

func (s *Bootstrap) loadRegistry() {
	s.r = registry.NewZookeeperRegistry(strings.Split(cfg.C.ZkAddr, ","))
}

func (s *Bootstrap) loadElection() {
	s.electionSignal = make(chan bool, 1)
	s.electionService = election.NewElection(s.electionSignal, s.r)
	s.c = cluster.NewCluster(s.electionSignal, s.w)
}

func (s *Bootstrap) loadWorker() {
	tsk_type := cfg.C.TaskLoadType
	if strings.EqualFold(tsk_type, "local") {
		s.w = worker.NewWorkerWithCfg(s.r)
	} else {
		s.w = worker.NewWorkerWithCfg(s.r)
	}
}
