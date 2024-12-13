package controllers

type AuthController struct {
	BaseController
	Authed bool
	Username string
}

func (c *AuthController) Prepare() {
	c.BaseController.Prepare()
	username := c.GetSession("username")
	c.Authed = username != nil
	if c.Authed {
		c.Username = username.(string)
	}
	c.Data["Authed"] = c.Authed
	c.Data["Username"] = c.Username
}

func (c *AuthController) NeedLogin() bool {
	if !c.Authed {
		c.Redirect("/login", 302)
		c.StopRun()
	}
	return !c.Authed
}
