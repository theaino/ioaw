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

func (c *ArticleController) CreateForm() {
	if c.NeedLogin() { return }
	c.TplName = "article/new.html"
}

func (c *ArticleController) Create() {
	if c.NeedLogin() { return }

	title := c.GetString("title")
	description := c.GetString("description")
	body := c.GetString("body")
	article := models.Article{
		Title: title,
		Description: description,
		Body: body,
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

func (c *ArticleController) EditForm() {
    id, _ := c.GetInt(":id")
    o := orm.NewOrm()
    article := models.Article{Id: id}
    err := o.Read(&article)
    if err == nil {
        c.Data["Article"] = article
        c.TplName = "article/edit.html"
    } else {
        c.Redirect("/articles", 302)
    }
}

func (c *ArticleController) Update() {
    id, _ := c.GetInt(":id")
    o := orm.NewOrm()
    article := models.Article{Id: id}
    if o.Read(&article) == nil {
        article.Title = c.GetString("title")
        article.Description = c.GetString("description")
        article.Body = c.GetString("body")
        o.Update(&article)
    }
    c.Redirect("/articles", 302)
}

func (c *ArticleController) Delete() {
    id, _ := c.GetInt(":id")
    o := orm.NewOrm()
    o.Delete(&models.Article{Id: id})
    c.Redirect("/articles", 302)
}
