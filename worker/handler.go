package worker

/*
 * Copyright 2020-2021 the original author(https://github.com/wj596)
 *
 * <p>
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * </p>
 */

import (
	"encoding/json"
	"fmt"
	"github.com/l-angel/tunnel/cfg"
	"github.com/l-angel/tunnel/meta/schema"
	"github.com/l-angel/tunnel/worker/sink"
	"log"
	"time"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
)

type handler struct {
	id    string
	queue chan interface{}
	stop  chan struct{}
	s     sink.Sink
	store *Store
	c     *cfg.CfgTask
}

func newHandler(taskId string, c *cfg.CfgTask, s sink.Sink, store *Store) *handler {

	return &handler{
		id:    taskId,
		queue: make(chan interface{}, 4096),
		stop:  make(chan struct{}, 1),
		s:     s,
		c:     c,
		store: store,
	}
}

func (s *handler) OnRotate(e *replication.RotateEvent) error {
	s.queue <- NewPositionWithForce(string(e.NextLogName), e.Position)
	return nil
}

func (s *handler) OnTableChanged(schema, table string) error {
	//err := _transferService.updateRule(schema, table)
	//if err != nil {
	//	return errors.Trace(err)
	//}
	return nil
}

func (s *handler) OnDDL(nextPos mysql.Position, _ *replication.QueryEvent) error {
	s.queue <- NewPositionWithForce(nextPos.Name, uint64(nextPos.Pos))
	return nil
}

func (s *handler) OnXID(nextPos mysql.Position) error {
	s.queue <- NewPosition(nextPos.Name, uint64(nextPos.Pos), false)
	return nil
}

func (s *handler) OnRow(event *canal.RowsEvent) error {
	e := OEvent{Action: Of(event.Action), Rows: event.Rows, Table: event.Table.Name, Columns: schema.FromColumns(event.Table.Columns)}
	var mapping Mapping
	if e.Action == Update {
		mapping = &UpdateMapping{}
	} else if e.Action == Insert {
		mapping = &InsertMapping{}
	} else if e.Action == Delete {
		mapping = &DeleteMapping{}
	} else {
		//Todo only support insert,update,delete action.
		return nil
	}
	var requests []interface{}
	requests = append(requests, mapping.mapTo(e))
	s.queue <- requests

	return nil
}

func (s *handler) OnGTID(gtid mysql.GTIDSet) error {
	return nil
}

func (s *handler) OnPosSynced(pos mysql.Position, set mysql.GTIDSet, force bool) error {
	s.queue <- NewPositionWithForce(pos.Name, uint64(pos.Pos))

	return nil
}

func (s *handler) String() string {
	return "CanalEventHandler"
}

func (s *handler) startListener() {
	go func() {
		interval := time.Duration(s.c.FlushInterval)
		if interval <= 0 {
			interval = 5
		}
		bufferSize := s.c.FlushBufferSize

		ticker := time.NewTicker(time.Millisecond * interval)
		defer ticker.Stop()

		lastSavedTime := time.Now()
		requests := make([]interface{}, 0, bufferSize)
		var current *Position
		from := &Position{}
		d := s.store.get()
		err := json.Unmarshal(d, from)
		if err != nil {
			// todo stop task
		}
		fmt.Println(from)
		for {
			needFlush := false
			needSavePos := false
			select {
			case v := <-s.queue:
				switch v := v.(type) {
				case *Position:
					now := time.Now()
					if v.Force || now.Sub(lastSavedTime) > 3*time.Second {
						lastSavedTime = now
						needFlush = true
						needSavePos = true
						current = v
					}
				case []interface{}:
					requests = append(requests, v...)
					needFlush = int64(len(requests)) >= bufferSize
				}
			case <-ticker.C:
				needFlush = true
			case <-s.stop:
				return
			}

			if needFlush && len(requests) > 0 {
				s.sinkAsync(requests)
				requests = requests[0:0]
			}
			if needSavePos {
				d, _ := json.Marshal(Position{
					Name: current.Name,
					Pos:  current.Pos,
				})
				err := s.store.save(d)

				if err != nil {
					//todo stop task
				}
				from = current
			}
		}
	}()
}

func (s *handler) sinkAsync(v interface{}) {
	go func(d interface{}) {
		if s.s != nil {
			s.s.Sink(v)
		}
	}(v)
}
func (s *handler) stopListener() {
	log.Println("transfer stop")
	s.stop <- struct{}{}
}
