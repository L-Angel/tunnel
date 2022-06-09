package worker

import (
	"encoding/json"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/l-angel/tunnel/cfg"
	"github.com/l-angel/tunnel/log"
	"github.com/l-angel/tunnel/registry"
	sink2 "github.com/l-angel/tunnel/worker/sink"
	"go.uber.org/atomic"
	"strconv"
	"sync"
)

type EventListener interface {
	on()
}

type TaskStatus int32

const (
	Waiting TaskStatus = 0 // 未初始化
	Running TaskStatus = 1 // 工作中
	Pending TaskStatus = 2 // 已初始化，未开始运行
)

type Task struct {
	id   string
	lock sync.Mutex

	stopSignal   chan bool
	canal        *canal.Canal
	canalConfig  *canal.Config
	canalHandler *handler
	initialized  atomic.Bool

	//myId        string
	myAddr     string
	myPort     int
	myUsername string
	myPassword string
	myPosition *Position
	schemas    []*Schema

	rules  []string
	tskCfg *cfg.CfgTask
	sink   sink2.Sink

	Status TaskStatus

	store *Store
}

type Schema struct {
	dbName string
	tables []string
}

func NewTask(id string, t *cfg.CfgTask, r registry.Registry) *Task {
	var schemas []*Schema
	for _, cs := range t.Schemas {
		var tables []string
		for _, cst := range cs.Tables {
			tables = append(tables, cst.Name)
		}
		schemas = append(schemas, &Schema{
			dbName: cs.Name,
			tables: tables,
		})
	}

	var rules []string
	for _, s := range schemas {
		if len(s.tables) <= 0 {
			rules = append(rules, s.dbName+".*\\..*")
		} else {
			for _, t := range s.tables {
				rules = append(rules, s.dbName+".*\\."+t+".*")
			}
		}
	}

	return &Task{
		id:         id,
		myAddr:     t.Addr,
		myPort:     t.Port,
		myUsername: t.Username,
		myPassword: t.Password,
		tskCfg:     t,
		schemas:    schemas,
		rules:      rules,
		sink:       toSink(id, t.Sink),
		store:      newStore(id, r),
	}
}

func toSink(taskId string, sc map[string]interface{}) sink2.Sink {
	if sc["type"] == "file" {
		return sink2.NewSink(taskId, sink2.NewCfgFileSink(sc["path"].(string)))
	}
	return sink2.NewSink(taskId, sink2.NewCfgStdSink())
}

func (t *Task) StartUp() error {
	t.lock.Lock()
	defer t.lock.Unlock()

	if !t.initialized.Load() {
		t.Initialize()
		t.initialized.Store(true)
		if err := t.createCanal(); err != nil {
			return err
		}
		t.canalHandler = newHandler(t.id, t.tskCfg, t.sink, t.store)
		t.canal.SetEventHandler(t.canalHandler)
		//t.sink.Initialize()
		t.canalHandler.startListener()
		t.Run()
	} else {
		return t.restart()
	}
	return nil
}

func (t *Task) restart() error {
	if err := t.createCanal(); err != nil {
		return err
	}
	t.canalHandler = newHandler(t.id, t.tskCfg, t.sink, t.store)
	t.canal.SetEventHandler(t.canalHandler)
	t.sink.Initialize()
	t.canalHandler.startListener()
	t.Run()
	return nil
}

func (t *Task) Initialize() {
	t.canalConfig = canal.NewDefaultConfig()
	t.canalConfig.Addr = t.myAddr + ":" + strconv.Itoa(t.myPort)
	t.canalConfig.Password = t.myPassword
	t.canalConfig.User = t.myUsername
	t.canalConfig.Dump = canal.DumpConfig{}

	t.canalConfig.IncludeTableRegex = t.rules
	t.canalConfig.ExcludeTableRegex = []string{".*\\.__drds_.*"}

	if t.store != nil {
		pos := Position{}
		d := t.store.get()
		err := json.Unmarshal(d, &pos)
		if err != nil {
			t.myPosition = NewPositionWithForce(pos.Name, pos.Pos)
		}
	}
	t.Status = Pending
	t.sink.Initialize()
}

func (t *Task) createCanal() error {
	var err error
	t.canal, err = canal.NewCanal(t.canalConfig)
	return err
}

func (t *Task) Run() {
	if t.canal == nil {
		return
	}
	go func(cana *canal.Canal, pos *Position) {
		var err error
		if pos == nil {
			err = cana.Run()
		} else {
			err = cana.RunFrom(mysql.Position{pos.Name, uint32(pos.Pos)})
		}
		if err != nil {
			// todo running set true
			log.Error(err)
		} else {
			t.Status = Running
		}

	}(t.canal, t.myPosition)
}
func (t *Task) listen() {

}
func (t *Task) Stop() {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.Status = Waiting
	if t.canalHandler != nil {
		t.canalHandler.stopListener()
		t.canalHandler = nil
	}
	if t.canal != nil {
		t.canal.Close()
		t.canal = nil
	}
	if t.sink != nil {
		t.sink.Close()
	}
}
