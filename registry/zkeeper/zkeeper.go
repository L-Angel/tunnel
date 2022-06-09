package zkeeper

import (
	"github.com/samuel/go-zookeeper/zk"
)

type KeepWatcher struct {
	conn *zk.Conn
}

func NewKeepWatcher(conn *zk.Conn) *KeepWatcher {
	return &KeepWatcher{conn: conn}
}

func (kw *KeepWatcher) WatchData(path string, listener func(data []byte, err error)) {
	go func(path string, w *KeepWatcher) {
		for {
			data, _, event, err := w.conn.GetW(path)
			listener(data, err)
			<-event
		}
	}(path, kw)
}

func (kw *KeepWatcher) WatchChildren(path string, listener func(children []string, eventType zk.Event, err error)) {
	go func(path string, w *KeepWatcher) {
		for {
			children, _, event, err := w.conn.ChildrenW(path)
			eventType := <-event
			listener(children, eventType, err)
		}
	}(path, kw)
}
