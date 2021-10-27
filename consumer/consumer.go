package consumer

import (
	"context"
	"fmt"
	"go-web/conf"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func Start() {
	con, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName(conf.Config.RocketMQ.Consumer),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{conf.Config.RocketMQ.NameSrv})),
	)
	con.Subscribe(
		conf.Config.RocketMQ.Topic,
		consumer.MessageSelector{},
		func(c context.Context, me ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range me {
				fmt.Println(string(me[i].Message.Body))
				// buf := bytes.NewBuffer(me[i].Message.Body)
				// var rece message.Cpu
				// err := binary.Read(buf, binary.LittleEndian, &rece)
				// if err != nil {
				// 	fmt.Println("consumer err,err is ", err)
				// 	os.Exit(1)
				// }
				// message.Consumer(&rece)
			}
			return consumer.ConsumeSuccess, nil
		},
	)
	con.Start()
	for {
	}
}
