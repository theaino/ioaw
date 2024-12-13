package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Article struct {
	Id       int
	Title string
	Description string
	Body string
}

func init() {
	orm.RegisterModel(new(Article))
}
