package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"ioaw/models"

	"github.com/beego/beego/v2/client/orm"
)

type UserController struct {
	AuthController
}

func (c *UserController) CheckUserExist() bool {
	o := orm.NewOrm()
	userCount, err := o.QueryTable(new(models.User)).Count()
	if userCount != 0 || err != nil {
		c.Redirect("/login", 302)
		return true
	}
	return false
}

func (c *UserController) RegisterForm() {
	if c.CheckUserExist() { return }

	o := orm.NewOrm()
	userCount, err := o.QueryTable(new(models.User)).Count()
	if userCount != 0 || err != nil {
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "user/register.html"
}

func (c *UserController) Register() {
	if c.CheckUserExist() { return }
	
	username := c.GetString("username")
	password := c.GetString("password")

	hasher := md5.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	user := models.User{
		Username: username,
		Password: hashedPassword,
	}

	o := orm.NewOrm()
	_, err := o.Insert(&user)
	if err != nil {
		c.Data["Error"] = "User already exists"
		c.TplName = "user/register.html"
		return
	}

	c.Redirect("/login", 302)
}

func (c *UserController) LoginForm() {
	c.TplName = "user/login.html"
}

func (c *UserController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")

	hasher := md5.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	user := models.User{Username: username}
	o := orm.NewOrm()
	err := o.Read(&user, "Username")

	if err == nil && user.Password == hashedPassword {
		c.SetSession("username", username)
		c.Redirect("/articles", 302)
	} else {
		c.Data["Error"] = "Invalid username or password"
		c.TplName = "user/login.html"
	}
}

func (c *UserController) Logout() {
	if c.Authed {
		c.DestroySession()
	}
	c.Redirect("/login", 302)
}
