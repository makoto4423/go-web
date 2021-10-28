package consumer

import (
	"context"
	_ "encoding/json"
	"fmt"
	"go-web/conf"
	_ "go-web/message"
	_ "log"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/golang/protobuf/proto"

	protoCompany "go-web/proto/company"
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
			arr := make([]byte, 0)
			for i := range me {
				arr = append(arr, me[i].Message.Body...)
				// buf := bytes.NewBuffer(me[i].Message.Body)
				// var rece message.Cpu
				// err := binary.Read(buf, binary.LittleEndian, &rece)
				// if err != nil {
				// 	fmt.Println("consumer err,err is ", err)
				// 	os.Exit(1)
				// }
				// message.Consumer(&rece)
			}
			// rece := message.Message{}
			// json.Unmarshal(arr, &rece)
			// message.Consumer(&rece)
			// fmt.Println(rece)
			company := protoCompany.Company{}
			proto.Unmarshal(arr, &company)
			fmt.Println("company ", company.Name)
			return consumer.ConsumeSuccess, nil
		},
	)
	con.Start()
	for {
	}
}
