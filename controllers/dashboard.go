package controllers

type DashboardController struct {
	AuthController
}

func (c *DashboardController) Index() {
	c.TplName = "index.html"
}
