package controllers

import (
	"ioaw/models"

	"github.com/beego/beego/v2/client/orm"
)

type ArticleController struct {
	AuthController
}

func (c *ArticleController) List() {
	o := orm.NewOrm()
	var articles []models.Article
	o.QueryTable(new(models.Article)).All(&articles)
	c.Data["Articles"] = articles
	c.TplName = "article/list.html"
}

func (c *ArticleController) View() {
	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err := o.Read(&article)
	if err == nil {
		c.Data["Article"] = article
		c.TplName = "article/view.html"
	} else {
		c.Redirect("/articles", 302)
	}
}

func (c *ArticleController) New() {
	if c.NeedLogin() { return }

	c.Data["Form"] = &models.Article{}
	c.TplName = "article/new.html"
}

func (c *ArticleController) Create() {
	if c.NeedLogin() { return }

	var article models.Article
	if c.ParseForm(&article) != nil {
		c.Data["Error"] = "Failed to create article."
		c.TplName = "article/create.html"
		return
	}
	o := orm.NewOrm()
	_, err := o.Insert(&article)
	if err == nil {
		c.Redirect("/articles", 302)
	} else {
		c.Data["Error"] = "Failed to create article."
		c.TplName = "article/create.html"
	}
}

func (c *ArticleController) Edit() {
	if c.NeedLogin() { return }

	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err := o.Read(&article)
	if err == nil {
		c.Data["Article"] = article
		c.Data["Form"] = &article
		c.TplName = "article/edit.html"
	} else {
		c.Redirect("/articles", 302)
	}
}

func (c *ArticleController) Update() {
	if c.NeedLogin() { return }

	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	article := models.Article{Id: id}
	if o.Read(&article) == nil {
		c.ParseForm(&article)
		o.Update(&article)
	}
	c.Redirect("/articles", 302)
}

func (c *ArticleController) Delete() {
	if c.NeedLogin() { return }

	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	o.Delete(&models.Article{Id: id})
	c.Redirect("/articles", 302)
}
