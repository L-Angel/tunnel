package sink

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	defaultFileName = "data.seg"
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
	self.file, _ = os.OpenFile(self.path+"/"+self.name+"_"+self.taskId, os.O_APPEND|os.O_RDWR, 0666)
}

func (self *FileSink) check() {
	if self.path == "" {
		panic(fmt.Sprintf("Check sink file error, is empty!"))
	}
	_, err := os.Open(self.path + "/" + self.name)
	if err != nil && os.IsNotExist(err) {
		_ = os.MkdirAll(self.path, os.ModePerm)
		_, _ = os.Create(self.path + "/" + self.name)
	}
}

func (self *FileSink) Sink(v interface{}) {
	r, _ := json.Marshal(v)
	w := bufio.NewWriter(self.file)
	_, _ = w.WriteString(strconv.FormatInt(time.Now().UnixNano(), 10) + "|" + string(r) + "\n")
	_ = w.Flush()
}

func (self *FileSink) Close() {
	if self.file == nil {
		return
	}
	_ = self.file.Close()
}
