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
	f, _ := engine.Where(user.Id == 1).Get(user)
	if !f {
		return nil
	}
	return user
}

func List() []entity.User {
	list := make([]entity.User, 0)
	user := entity.User{}
	_ = engine.Where(user.Id < 10 && user.Id > 5).Limit(1, 2).Find(&list)
	return list
}
