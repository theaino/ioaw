package controllers

import "github.com/beego/beego/v2/server/web"

type BaseController struct {
	web.Controller
}

func (c *BaseController) Prepare() {
	c.Layout = "layout/base.html"
}
