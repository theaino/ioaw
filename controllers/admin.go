package controllers

type AdminController struct {
	AuthController
}

func (c *AdminController) Prepare() {
	c.AuthController.Prepare()
	c.NeedLogin()
}
