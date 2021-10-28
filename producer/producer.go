package producer

import (
	"context"
	"fmt"
	"go-web/conf"
	_ "go-web/message"
	_ "math/rand"
	"time"

	protoCompany "go-web/proto/company"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/golang/protobuf/proto"
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
		// p.SendSync(context.Background(), primitive.NewMessage(conf.Config.RocketMQ.Topic, message.Producer(rand.Int()%3)))
		company := protoCompany.Company{
			Name: "web",
			Code: 4,
		}
		arr, err := proto.Marshal(&company)
		if err != nil {
			fmt.Println(err)
		}
		p.SendSync(context.Background(), primitive.NewMessage(conf.Config.RocketMQ.Topic, arr))
		time.Sleep(time.Minute)
	}

}
