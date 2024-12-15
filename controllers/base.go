package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

type BaseController struct {
	web.Controller
	i18n.Locale
}

func (c *BaseController) Prepare() {
	c.Layout = "layout/base.html"

	lang := c.GetString("lang")
	if lang == "" {
		lang = c.Ctx.Input.Header("Accept-Language")
	}

	if !i18n.IsExist(lang) {
		lang = "en-US"
	}

	c.Lang = lang
	c.Data["Lang"] = lang
}

