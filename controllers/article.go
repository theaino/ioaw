package controllers

import (
	"html/template"
	"ioaw/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
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
	if err != nil {
		c.Redirect(web.URLFor("ArticleController.List"), 302)
		return
	}
	content, err := article.Render()
	if err != nil {
		logs.Error(err)
		c.Redirect(web.URLFor("ArticleController.List"), 302)
		return
	}
	c.Data["Title"] = article.Title
	c.Data["Article"] = article
	c.Data["Content"] = template.HTML(content)
	c.TplName = "article/view.html"
}

func (c *ArticleController) New() {
	if c.NeedLogin() { return }

	c.Data["Article"] = models.Article{}
	c.TplName = "article/new.html"
}

func (c *ArticleController) Create() {
	if c.NeedLogin() { return }

	var article models.Article
	if err := c.ParseForm(&article); err != nil {
		logs.Error("Failed to create article", err)
		c.Data["Error"] = "Failed to create article."
		c.TplName = "article/new.html"
		return
	}

	if err := article.Parse(article.Body); err != nil {
		c.Data["Error"] = c.Tr("article.error_body_invalid")
		c.TplName = "article/new.html"
		return
	}

	o := orm.NewOrm()
	_, err := o.Insert(&article)
	if err == nil {
		c.Redirect(web.URLFor("ArticleController.List"), 302)
	} else {
		logs.Error("Failed to create article", err)
		c.Data["Error"] = "Failed to create article."
		c.TplName = "article/new.html"
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
		c.TplName = "article/edit.html"
	} else {
		c.Redirect(web.URLFor("ArticleController.List"), 302)
	}
}

func (c *ArticleController) Update() {
	if c.NeedLogin() { return }

	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	article := models.Article{Id: id}
	if o.Read(&article) != nil {
		c.Redirect(web.URLFor("ArticleController.List"), 302)
		return
	}
	c.ParseForm(&article)
	if err := article.Parse(article.Body); err != nil {
		c.Data["Error"] = c.Tr("article.error_body_invalid")
		c.TplName = "article/new.html"
		return
	}
	o.Update(&article)
	c.Redirect(web.URLFor("ArticleController.View", ":id", article.Id), 302)
}

func (c *ArticleController) Delete() {
	if c.NeedLogin() { return }

	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	o.Delete(&models.Article{Id: id})
	c.Redirect(web.URLFor("ArticleController.List"), 302)
}
