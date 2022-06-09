package sink

import "time"

type Message struct {
	TraceId string    `json:"trace_id"`
	From    string    `json:"from"`
	Created time.Time `json:"created"`

	Headers []Header    `json:"headers"`
	Body    interface{} `json:"body"`
}

type Header struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}