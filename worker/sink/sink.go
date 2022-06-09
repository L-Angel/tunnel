package sink

import (
	"reflect"
	"strings"
)

const defaultTopic = "tunnel_sink_topic"

type CfgSink interface {
	Type() string
}

type Sink interface {
	Initialize()
	Sink(v interface{})
	Close()
}

func NewSink(taskId string, cs CfgSink) Sink {
	t := cs.Type()
	rto := reflect.TypeOf(cs)
	if strings.EqualFold(t, "file") && rto == reflect.TypeOf(&CfgFileSink{}) {
		return newFileSink(taskId, cs.(*CfgFileSink))
	}
	return newStdSink(cs.(*CfgStdSink))
}
