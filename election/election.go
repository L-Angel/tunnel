package election

import (
	"errors"
	"fmt"
	"github.com/l-angel/tunnel/cfg"
	"github.com/l-angel/tunnel/log"
	"github.com/l-angel/tunnel/meta/net"
	"github.com/l-angel/tunnel/registry"
	"go.uber.org/atomic"
	"strings"
	"sync"
)

type Election struct {
	r registry.Registry

	signal       chan bool
	electionPath string

	once sync.Once

	selected atomic.Bool
	leader   atomic.String

	connectionAmount atomic.Int32
	degraded         atomic.Bool
}

func NewElection(_signal chan bool, r registry.Registry) *Election {
	return &Election{
		r:            r,
		electionPath: strings.Join([]string{cfg.PRoot, cfg.C.Cluster, cfg.C.Group, "election"}, "/"),
		signal:       _signal,
	}
}

func (self *Election) IsLeader() bool {
	return self.selected.Load()
}

func (self *Election) beLeader() {
	self.selected.Store(true)
	self.leader.Store("self")
	self.signal <- self.selected.Load()
}

func (self *Election) beFollower(leader string) {
	self.leader.Store(leader)
	self.selected.Store(false)
	self.signal <- self.selected.Load()
}

func (self *Election) Elect() error {
	ip, _ := net.ExternalIP()
	err := self.r.Create(self.electionPath, true, []byte(ip.String()))
	if err == nil {
		self.beLeader()
	} else {
		d := self.r.GetData(self.electionPath)
		if d == nil {
			return errors.New("Get election data error.")
		}
		leader := string(d)
		ip, _ := net.ExternalIP()
		if leader == ip.String() {
			self.beLeader()
		} else {
			self.beFollower(leader)
		}
	}

	self.once.Do(func() {
		self.startWatchConnectingState()
		self.startWatchNodeState()
	})
	return nil
}

func (self *Election) startWatchConnectingState() {

	self.r.OnChange(func(eventType registry.EventType) {
		if self.selected.Load() {
			if eventType == registry.Connecting {
				self.connectionAmount.Inc()
			}
		}

		if eventType == registry.Connected {
			ip, _ := net.ExternalIP()
			err := self.r.Create(strings.Join([]string{cfg.PRoot, cfg.C.Cluster, "_node", ip.String()}, "/"), true, nil)
			fmt.Println(err)
			self.connectionAmount.Store(0)
			if self.degraded.Load() {
				_ = self.Elect()
				self.degraded.Store(false)
			}
		}
	})
}

func (self *Election) startWatchNodeState() {
	self.r.On(self.electionPath, func(data []byte, eventType registry.NodeEventType, err error) {
		if eventType == registry.NodeDelete {
			err := self.Elect()
			if err != nil {
				log.Errorf("elect new master error %s ", err.Error())
			}
		}
	})
}

func (self *Election) degrading() {
	if !self.degraded.Load() {
		self.degraded.Store(true)
		self.beFollower("")
	}
}
