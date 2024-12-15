package models

import (
	libarticle "ioaw/lib/article"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Article struct {
	Id       int `form:"-"`
	Title string `form:"-"`
	Description string `form:"-"`
	Time time.Time `form:"-"`
	Body string `form:"body"`
}

func init() {
	orm.RegisterModel(new(Article))
}

func (a *Article) Parse(raw string) error {
	doc, err := libarticle.Parse(raw, a.Title)
	if err != nil {
		return err
	}
	a.Title = doc.Title
	a.Description = doc.Summary
	a.Time = doc.Time
	a.Body = raw
	return nil
}

func (a *Article) Render() (string, error) {
	doc, err := libarticle.Parse(a.Body, a.Title)
	if err != nil { return "", err }
	return libarticle.Render(doc)
}
