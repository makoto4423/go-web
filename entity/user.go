package entity

import (
	"fmt"
	"go-web/orm"
	"log"
)

type User struct {
	//Id       int8   `json:"id" xorm:"int(11)"`
	//Name     string `json:"name" xorm:"varchar(50)"`
	//Password string `json:"password" xorm:"varchar(50)"`
	//Status   string `json:"status" xorm:"varchar(10)"`
	Id       int8   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   string `json:"status"`
}

func init() {
	test := orm.GetTest()
	fmt.Print(test)
	engine := orm.GetEngine()
	user := new(User)
	exists, err := engine.IsTableExist(user)
	if err != nil {
		log.Fatal("error")
	}
	if !exists {
		engine.CreateTables(user)
	}
}
