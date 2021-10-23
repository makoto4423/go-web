package orm

import (
	"go-web/conf"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
	test   *Test
)

type Test struct {
	val string
}

func init() {
	conf.Init()
	var err error
	/**
	engine, err := xorm.NewEngine("mysql", conf.Config.DataSource.Url)
	这样写会导致init里engine的赋值失效，就是调用GetEngine返回的是个nil，为啥，或许这就是指针吧orz
	*/
	engine, err = xorm.NewEngine("mysql", conf.Config.DataSource.Url)
	if err != nil {
		//log.Default().Fatal("init engine failed")
		log.Fatal("init engine failed")
		os.Exit(-1)
	}
	engine.SetMaxIdleConns(int(conf.Config.DataSource.IdleCon))
	engine.SetMaxOpenConns(int(conf.Config.DataSource.OpenCon))
	engine.ShowSQL(conf.Config.DataSource.ShowSQL)
	engine.ShowExecTime(conf.Config.DataSource.ShowExecTime)
	test = new(Test)
	test.val = "abc"
}

func GetEngine() *xorm.Engine {
	return engine
}

func GetTest() *Test {
	return test
}
