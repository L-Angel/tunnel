package sink

import (
	"encoding/json"
	"fmt"
)

type CfgStdSink struct {
	CfgSink
	t string `yaml:"type"`
}

func NewCfgStdSink() *CfgStdSink {
	return &CfgStdSink{t: "std"}
}
func (c *CfgStdSink) Type() string {
	return c.t
}

type StdSink struct {
}

func newStdSink(c *CfgStdSink) *StdSink {
	return &StdSink{}
}

func (self *StdSink) Initialize() {

}

func (self *StdSink) Sink(v interface{}) {
	r, _ := json.Marshal(v)
	fmt.Printf("%v\n", string(r))
}

func (self *StdSink) Close() {

}
