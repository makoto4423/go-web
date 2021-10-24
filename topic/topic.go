package topic

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"go-web/conf"
	"os"
)

func init() {
	nameSrv := primitive.NewPassthroughResolver([]string{conf.Config.RocketMQ.NameSrv})
	resolver := admin.WithResolver(nameSrv)
	a, err := admin.NewAdmin(resolver)
	if err != nil {
		os.Exit(-1)
	}
	err = a.CreateTopic(context.Background(), admin.WithTopicCreate(conf.Config.RocketMQ.Topic), admin.WithBrokerAddrCreate(conf.Config.RocketMQ.Broker))
	if err != nil {
		//log.Fatal(err)
		os.Exit(-1)
	}
}
