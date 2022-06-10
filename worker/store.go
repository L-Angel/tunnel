package worker

import (
	"github.com/l-angel/tunnel/cfg"
	"github.com/l-angel/tunnel/registry"
	"strings"
)

type Store struct {
	taskId string
	r      registry.Registry
	node   string
}

func newStore(taskId string, r registry.Registry) *Store {
	s := &Store{r: r, node: "/" + strings.Join([]string{cfg.PRoot, cfg.C.Cluster, "task", taskId, "position"}, "/")}
	_ = s.r.Create(s.node, false, nil)
	return s
}

func (s *Store) save(data []byte) error {
	return s.r.SetData(s.node, data)
}

func (s *Store) get() []byte {
	return s.r.GetData(s.node)
}
