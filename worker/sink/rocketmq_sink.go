package sink

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/l-angel/tunnel/log"
)

type RocketMqSink struct {
	Retry     int
	GroupName string
	Endpoint  []string
	producer  rocketmq.Producer
}

func (self *RocketMqSink) Initialize() {
	var err error
	self.producer, err = rocketmq.NewProducer(producer.WithNameServer(self.Endpoint),
		producer.WithNsResovler(primitive.NewPassthroughResolver(self.Endpoint)),
		producer.WithRetry(self.Retry),
		producer.WithGroupName(self.GroupName))

	if err != nil {
		log.Error(err)
	}
}

func (self *RocketMqSink) Sink(data interface{}) {
	r, err := json.Marshal(data)
	msg := primitive.Message{Topic: defaultTopic, Body: r}

	if err != nil {
		log.Error(err)
		return
	}
	err = self.producer.SendAsync(context.Background(), func(ctx context.Context, result *primitive.SendResult, err error) {
		if err != nil {
			log.Error(err)
		}
	}, &msg)

	if err != nil {
		log.Error(err)
	}
}
