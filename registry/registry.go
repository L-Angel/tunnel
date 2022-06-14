package registry

import (
	"github.com/l-angel/tunnel/meta/znodes"
	"github.com/l-angel/tunnel/registry/zkeeper"
	"github.com/samuel/go-zookeeper/zk"
	"go.uber.org/atomic"
	"time"
)

type EventType int32

const (
	Connecting EventType = 1
	Connected  EventType = 2
)

type NodeEventType int32

const (
	Create         NodeEventType = 1
	Delete         NodeEventType = 2
	DataChange     NodeEventType = 3
	ChildrenChange NodeEventType = 4
	NodeDelete     NodeEventType = 5
)

type Registry interface {
	Create(path string, ephemeral bool, data []byte) error
	Delete(path string) error

	SetData(path string, data []byte) error
	GetData(path string) []byte

	On(path string, listener func(data []byte, eventType NodeEventType, err error))
	OnChange(listener func(listener EventType))
}

type ZookeeperRegistry struct {
	conn             *zk.Conn
	servers          []string
	timeout          time.Duration
	signal           <-chan zk.Event
	keepWatcher      *zkeeper.KeepWatcher
	connectionAmount *atomic.Int32
}

func NewZookeeperRegistry(servers []string) *ZookeeperRegistry {
	zkRegistry := &ZookeeperRegistry{servers: servers}
	zkRegistry.connect()
	return zkRegistry
}

func (r *ZookeeperRegistry) connect() {
	r.conn, r.signal, _ = zk.Connect(r.servers, 30*time.Second)
	r.keepWatcher = zkeeper.NewKeepWatcher(r.conn)
}

func (r *ZookeeperRegistry) Create(path string, ephemeral bool, data []byte) error {
	return znodes.CreatePathWithDataIfNecessary(r.conn, path, ephemeral, data)
}

func (r *ZookeeperRegistry) Delete(path string) error {
	_, stat, _ := r.conn.Get(path)
	if stat == nil {
		return r.conn.Delete(path, 0)
	} else {
		return r.conn.Delete(path, stat.Version+1)
	}
}

func (r *ZookeeperRegistry) SetData(path string, data []byte) error {
	_, stat, _ := r.conn.Get(path)
	if stat == nil {
		_, err := r.conn.Set(path, data, 0)
		return err
	} else {
		_, err := r.conn.Set(path, data, stat.Version)
		return err
	}
}

func (r *ZookeeperRegistry) GetData(path string) []byte {
	d, _, _ := r.conn.Get(path)
	return d
}

func (r *ZookeeperRegistry) On(path string, listener func(data []byte, eventType NodeEventType, err error)) {
	r.keepWatcher.WatchData(path, func(data []byte, err error) {
		listener(data, DataChange, err)
	})

	r.keepWatcher.WatchChildren(path, func(children []string, event zk.Event, err error) {
		if event.Type == zk.EventNodeDeleted {
			listener(nil, NodeDelete, err)
		}
	})
}

func (r *ZookeeperRegistry) OnChange(listener func(eventType EventType)) {

	go func(rr *ZookeeperRegistry) {
		for {
			select {
			case e := <-rr.signal:
				if e.State == zk.StateConnecting {
					listener(Connecting)
				}
				if e.State == zk.StateHasSession {
					listener(Connected)
				}
			}
		}
	}(r)
}
