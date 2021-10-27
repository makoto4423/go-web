package producer

import (
	"context"
	"go-web/conf"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func Start() {
	addr, _ := primitive.NewNamesrvAddr(conf.Config.RocketMQ.NameSrv)
	p, _ := rocketmq.NewProducer(
		producer.WithGroupName(conf.Config.RocketMQ.Producer),
		producer.WithNameServer(addr),
		producer.WithCreateTopicKey(conf.Config.RocketMQ.Topic),
		producer.WithRetry(1),
	)
	p.Start()
	for {
		p.SendSync(context.Background(), primitive.NewMessage(conf.Config.RocketMQ.Topic, []byte("Hello")))
		time.Sleep(time.Minute)
	}
}
