package zookeeper

import (
	"fmt"
	"go-web/conf"
	"time"

	"github.com/go-zookeeper/zk"
)

var con *zk.Conn
var event <-chan zk.Event
var err error

func init() {
	con, event, err = zk.Connect([]string{conf.Config.Zookeeper.Url}, time.Duration(60*1000*1000))
	if err != nil {
		fmt.Println("zookeeper init " + err.Error())
	}
}

func Get() {
	path := "/abc"
	flag, _, err := con.Exists(path)
	if err != nil {
		fmt.Println("zookeeper exists error" + err.Error())
	} else if flag {
		bytes, stat, _ := con.Get("/abc")
		fmt.Println("val is " + string(bytes) + ", version is " + string(stat.Version))
	} else {
		fmt.Println(path + " doesn`t exits")
	}
}
