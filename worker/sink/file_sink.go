package sink

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/l-angel/tunnel/log"
	"os"
	"strconv"
	"time"
)

const (
	defaultFileName   = "data"
	defaultFileSuffix = ".seg"
)

type CfgFileSink struct {
	CfgSink
	t      string `yaml:"type"`
	path   string `yaml:"path"`
	taskId string
}

func NewCfgFileSink(p string) *CfgFileSink {
	return &CfgFileSink{t: "file", path: p}
}

func (c *CfgFileSink) Type() string {
	return c.t
}

func (c *CfgFileSink) Path() string {
	return c.path
}

type FileSink struct {
	path   string
	name   string
	file   *os.File
	taskId string
}

func newFileSink(taskId string, c *CfgFileSink) *FileSink {
	return &FileSink{taskId: taskId, path: c.Path(), name: defaultFileName}
}

func (self *FileSink) Initialize() {
	self.check()
	self.file, _ = os.OpenFile(self.path+"/"+self.name+"_"+self.taskId+defaultFileSuffix, os.O_APPEND|os.O_RDWR, 0666)
}

func (self *FileSink) check() {
	if self.path == "" {
		panic(fmt.Sprintf("Check sink file error, is empty!"))
	}
	_, err := os.Open(self.path + "/" + self.name + "_" + self.taskId + defaultFileSuffix)
	if err != nil && os.IsNotExist(err) {
		_ = os.MkdirAll(self.path, os.ModePerm)
		_, _ = os.Create(self.path + "/" + self.name + "_" + self.taskId + defaultFileSuffix)
	}
}

func (self *FileSink) Sink(v interface{}) {
	r, _ := json.Marshal(v)
	w := bufio.NewWriter(self.file)
	var err error
	_, err = w.WriteString(strconv.FormatInt(time.Now().UnixNano(), 10) + "|" + string(r) + "\n")
	if err != nil {
		log.Error(err)
	}
	err = w.Flush()
	if err != nil {
		log.Error(err)
	}
}

func (self *FileSink) Close() {
	if self.file == nil {
		return
	}
	_ = self.file.Close()
}
