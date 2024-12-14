package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Article struct {
	Id       int `form:"-"`
	Title string `form:"title"`
	Description string `form:"description"`
	Body string `form:"body"`
}

func init() {
	orm.RegisterModel(new(Article))
}
