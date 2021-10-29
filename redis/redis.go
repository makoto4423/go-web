package redis

import (
	"fmt"
	"go-web/conf"

	"github.com/gomodule/redigo/redis"
)

func Get() {
	con, _ := redis.DialURL("redis://" + conf.Config.Redis.Url)
	reply, _ := con.Do("keys", "*")
	arr := reply.([]interface{})
	for _, val := range arr {
		fmt.Println(string(val.([]byte)))
	}
	con.Send("get", "abc")
	reply, _ = con.Receive()
	fmt.Println(reply)
}
