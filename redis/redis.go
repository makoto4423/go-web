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
	// var err error
	// reply, err = con.Receive()
	// if err != nil {
	// 	fmt.Print("redis get error")
	// } else {
	// 	fmt.Println(reply)

	// }
}
