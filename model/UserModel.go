package model

import (
	"go-web/entity"
	"go-web/orm"

	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	engine = orm.GetEngine()
}

func Get() *entity.User {
	user := new(entity.User)
	engine.Get(user)
	return user
}
